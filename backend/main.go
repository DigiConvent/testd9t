package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DigiConvent/testd9t/cli"
	cli_helpers "github.com/DigiConvent/testd9t/cli/helpers"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/file_repo"
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
		db.DatabasePath = os.Getenv("DATABASE_PATH")
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

	router := gin.Default()

	router.NoRoute(handleFrontend())
	err := router.RunTLS(":"+os.Getenv("PORT"), "/home/testd9t/certs/fullchain.pem", "/home/testd9t/certs/privkey.pem")

	if err != nil {
		panic("failed to start server: " + err.Error())
	}
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
	verbose := actionsFlagSet.Bool("verbose", false, "Run more verbosely")
	forceFlag := actionsFlagSet.Bool("force", false, "Apply fixes upon a failure during the installation")
	helpFlag := actionsFlagSet.Bool("help", false, "Prints this help message")
	installFlag := actionsFlagSet.String("install", "", "Install ")
	migrateDBFlag := actionsFlagSet.Bool("migrate-db", false, "Migrate the database to something that is compatible with the current version")
	replaceWithFlag := actionsFlagSet.String("replace-with", "", "Replace with a specific version")
	runFlag := actionsFlagSet.Bool("run", false, "Deploy")
	resetDBFlag := actionsFlagSet.Bool("reset-db", false, "Reset the database")
	statusFlag := actionsFlagSet.Bool("status", false, "Prints the current status")
	versionsFlag := actionsFlagSet.Bool("versions", false, "List all available versions")
	listFlavoursFlag := actionsFlagSet.String("supported-flavours", "", "List supported flavours")

	actionsFlagSet.Parse(os.Args[1:])

	if *replaceWithFlag != "" {
		fmt.Println("--replace-with", *replaceWithFlag)

		fromVersion := sys_domain.VersionFromString(*replaceWithFlag)
		if fromVersion == nil {
			fmt.Println("Invalid version", *replaceWithFlag)
			os.Exit(1)
		}

		url := "https://github.com/DigiConvent/testd9t/releases/download/" + *replaceWithFlag + "/main"
		err := file_repo.NewRepoRemote().DownloadAsset(url, "/home/testd9t/backend/main")
		if err != nil {
			fmt.Println("Error downloading new version:", err)
			os.Exit(1)
		}

		err = exec.Command("chmod", "+x", "/home/testd9t/backend/main").Run()
		if err != nil {
			fmt.Println("Error setting permissions:", err)
			os.Exit(1)
		}
	}

	if *resetDBFlag {
		cli.ResetDB()
	}

	if *helpFlag || actionsFlagSet.NFlag() == 0 {
		actionsFlagSet.Usage()
		os.Exit(0)
	}

	if *versionsFlag {
		cli_helpers.ListVersions(sysService)
	}

	if *listFlavoursFlag != "" {
		cli.ListFlavours(sysService)
	}

	if *statusFlag {
		cli.ShowStatus(sysService)
	}

	if *installFlag != "" {

		cli.Install(sysService, installFlag, *forceFlag, *verbose)
	}

	if *migrateDBFlag {
		cli.MigrateDB(sysService)
	}

	if !*runFlag {
		os.Exit(0)
	}
}
