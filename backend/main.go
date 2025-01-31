package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"syscall"

	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/install"
	services "github.com/DigiConvent/testd9t/pkg"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if sys_domain.ProgramVersion == "dev" {
		godotenv.Load(".env")
	} else {
		godotenv.Load("/home/digiconvent/env")
	}

	services := services.InitiateServices()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived shutdown signal, closing DB...")
		db.CloseAllDatabases()
		os.Exit(0)
	}()

	handleFlags(services.SysService)

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.NoRoute(handleFrontend())
	router.Run(":" + os.Getenv("PORT"))
}

func handleFrontend() gin.HandlerFunc {
	if sys_domain.ProgramVersion != "dev" {
		return func(c *gin.Context) {
			path := c.Request.URL.Path
			if path == "/favicon.ico" {
				path = "/favicon.ico"
			} else if strings.HasPrefix(path, "/assets/") {

			} else {
				path = "/index.html"
			}
			c.File(os.Getenv("STATIC_FILES") + "frontend/" + path)
		}
	} else {
		return proxyHandler("http://localhost:5173") // vite server proxy
	}
}

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			log.Fatalf("Could not parse proxy target URL: %v", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.RawQuery = c.Request.URL.RawQuery
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func handleFlags(sysService sys_service.SysServiceInterface) {
	actionsFlagSet := flag.NewFlagSet("Options", flag.ExitOnError)
	clearCacheFlag := actionsFlagSet.Bool("clear-cache", false, "Clear the cache")
	forceFlag := actionsFlagSet.Bool("force", false, "Apply fixes upon a failure during the installation")
	helpFlag := actionsFlagSet.Bool("help", false, "Prints this help message")
	installFlag := actionsFlagSet.String("install", "", "Install")
	// localFlag := actionsFlagSet.Bool("local", false, "Use local files")
	migrateDBFlag := actionsFlagSet.Bool("migrate-db", false, "Migrate the database to something that is compatible with the current version")
	runFlag := actionsFlagSet.Bool("run", false, "Deploy")
	resetDBFlag := actionsFlagSet.Bool("reset-db", false, "Reset the database")
	statusFlag := actionsFlagSet.Bool("status", false, "Prints the current status")
	versionsFlag := actionsFlagSet.Bool("versions", false, "List all available versions")
	versionFlavoursFlag := actionsFlagSet.String("supported-flavours", "", "List supported flavours for this version")

	actionsFlagSet.Parse(os.Args[1:])

	if *resetDBFlag {
		fmt.Println("--reset-db")

		dbPath := db.DatabasePath
		entries, err := os.ReadDir(dbPath)
		if err != nil {
			fmt.Println("Could not find db directory", dbPath)
		}

		for _, entry := range entries {
			dbName := entry.Name()
			if !entry.IsDir() {
				continue
			}

			dbPath := path.Join(dbPath, dbName+".db")
			err := os.RemoveAll(dbPath)
			if err != nil {
				fmt.Println("Error removing db:", err)
			}

			fmt.Println("Removed", dbPath)
		}
		os.Exit(0)
	}

	if *helpFlag || actionsFlagSet.NFlag() == 0 {
		actionsFlagSet.Usage()
		os.Exit(0)
	}

	sysStatus, _ := sysService.GetSystemStatus()

	if *versionsFlag {
		versions, status := sysService.ListReleaseTags()
		if status.Err() {
			fmt.Println("Error fetching versions:", status.Message)
			os.Exit(1)
		}

		if status.Err() {
			fmt.Println("Error fetching system status:", status.Message)
			os.Exit(1)
		}

		if sysStatus.ProgramVersion.Major == -1 {
			fmt.Println("Program version: dev")
			if sysStatus.DatabaseVersion.Equals(&sysStatus.ProgramVersion) {
				fmt.Println("No migrations found, run with --migrate-db to migrate to a compatible version")
			}
		}

		fmt.Println("┏━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━┓")
		fmt.Println("┃ Version ┃ Program ┃ Migration ┃")

		for _, tag := range versions {
			var program, migrationExecuted string
			if tag.Tag == sysStatus.ProgramVersion.String() {
				program = "✓"
			}
			if tag.Tag == sysStatus.DatabaseVersion.String() {
				migrationExecuted = "✓"
			}
			fmt.Println("┣━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━━━┫")
			fmt.Printf("┃ %7s ┃ %7s ┃ %9s ┃\n", tag.Tag, program, migrationExecuted)
		}
		fmt.Println("┗━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━━━┛")
	}

	if *versionFlavoursFlag != "" {
		install.GetFlavours(&sysStatus.ProgramVersion)
	}

	if *statusFlag {
		fmt.Println("--status")
		if sys_domain.ProgramVersion == "dev" {
			fmt.Println("Development")
		} else {
			fmt.Println("Production")
		}

		currentVersion := &sysStatus.DatabaseVersion
		if currentVersion == nil {
			fmt.Println("Migration  : No version found (try running with --migrate-db to migrate to something compatible with this version)")
		} else {
			fmt.Println("Migration  :", currentVersion.String())
		}
		fmt.Println("Environment:", sysStatus.ProgramVersion.String())

		showPackages(sysService)
	}

	if *installFlag != "" {
		uid := os.Geteuid()

		if uid != 0 {
			fmt.Println("You need to be root to install")
			os.Exit(1)
		}

		*installFlag = strings.ToLower(*installFlag)
		fmt.Println("--install")

		flavours := install.GetFlavours(&sysStatus.ProgramVersion)

		found := false
		for _, flavour := range flavours {
			if flavour == *installFlag {
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Flavour", *installFlag, "not found")
			choices := strings.Join(flavours, ", ")
			fmt.Println("Available flavours:", choices)
			os.Exit(1)
		}

		install.Install(installFlag, *forceFlag, *clearCacheFlag)
	}

	if *migrateDBFlag {
		fmt.Println("--migrate-db")

		packages := db.ListPackages()

		cols := []int{9, 13, 9, 8}
		fmt.Printf("┏%s┳%s┳%s┳%s┓\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
		fmt.Println("┃ Package ┃ Initialised ┃ Version ┃ Max V. ┃")

		for _, pkg := range packages {
			status := sysService.MigratePackage(pkg, sysStatus.ProgramVersion)
			if status.Err() && status.Code != 404 {
				fmt.Println("Error migrating package", pkg, ":", status.Message)
			}
		}

		fmt.Printf("┗%s┻%s┻%s┻%s┛\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
	}

	if !*runFlag {
		os.Exit(0)
	}
}

func showPackages(sysService sys_service.SysServiceInterface) {

	dbPackages, _ := sysService.GetPackages()
	packages := db.ListPackages()

	cols := []int{9, 13, 9, 8}
	fmt.Printf("┏%s┳%s┳%s┳%s┓\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
	fmt.Println("┃ Package ┃ Initialised ┃ Version ┃ Max V. ┃")

	for _, pkg := range packages {
		fmt.Printf("┣%s╋%s╋%s╋%s┫\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
		format := "┃ %" + strconv.Itoa(cols[0]-2) + "s ┃ %" + strconv.Itoa(cols[1]-2) + "s ┃ %" + strconv.Itoa(cols[2]-2) + "s ┃ %" + strconv.Itoa(cols[3]-2) + "s ┃\n"
		val, ok := dbPackages[pkg]
		initialised := "✓"
		var version, maxVersion string
		versions, _ := sysService.GetPackageVersions(pkg)
		sys_domain.Sort(versions, true)
		if len(versions) > 0 {
			maxVersion = versions[len(versions)-1].String()
		}
		if !ok {
			initialised = "✗"
		} else {
			version = val.Version.String()
		}
		fmt.Printf(format, pkg, initialised, version, maxVersion)
	}

	fmt.Printf("┗%s┻%s┻%s┻%s┛\n", strings.Repeat("━", cols[0]), strings.Repeat("━", cols[1]), strings.Repeat("━", cols[2]), strings.Repeat("━", cols[3]))
}
