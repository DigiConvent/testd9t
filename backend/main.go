package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/DigiConvent/testd9t/cli"
	cli_helpers "github.com/DigiConvent/testd9t/cli/helpers"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	"github.com/DigiConvent/testd9t/delivery/api"
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
		gin.SetMode(gin.ReleaseMode)
		log.SetLogLevel(0)
		db.DatabasePath = os.Getenv("DATABASE_PATH")
		err := godotenv.Load("/home/testd9t/env")
		if err != nil {
			log.Error(err.Error())
		}
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

	httpsRouter := gin.Default()
	httpsRouter.RedirectTrailingSlash = true

	if sys_domain.ProgramVersion == "dev" {
		log.Info("Development mode")
		httpsRouter.Use(gin.Logger())
		httpsRouter.Use(gin.Recovery())
		httpsRouter.Use(func(ctx *gin.Context) {
			log.Info("Request URL:" + ctx.Request.URL.Scheme)
			ctx.Next()
		})
		api.RegisterRoutes(httpsRouter, services)
		httpsRouter.NoRoute(handleFrontend())
		go func() {
			err := httpsRouter.Run(":8080")
			if err != nil {
				panic("failed to start dev server: " + err.Error())
			}
		}()
	} else {
		log.Info("Starting build from " + sys_domain.CompiledAt)
		httpsRouter.Use(func(ctx *gin.Context) {
			fmt.Println("Request URL:", ctx.Request.URL)
			if ctx.Request.URL.Scheme == "https" {
				ctx.Next()
				return
			}

			ctx.Redirect(http.StatusMovedPermanently, "https://"+ctx.Request.Host+ctx.Request.RequestURI)
		})
		api.RegisterRoutes(httpsRouter, services)
		httpsRouter.NoRoute(handleFrontend())

		go func() {
			err := httpsRouter.RunTLS(":"+os.Getenv("PORT"), "/home/testd9t/certs/fullchain.pem", "/home/testd9t/certs/privkey.pem")
			if err != nil {
				panic("failed to start server: " + err.Error())
			}
		}()
	}

	<-sigChan
	log.Info("Closing DigiConvent")
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
			c.File("frontend/" + path)
		}
	} else {
		return proxyHandler("http://localhost:5173")
	}
}

func proxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			log.Error(err.Error())
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
	logLevelFlag := actionsFlagSet.Int("log-level", 2, "Set the log level")

	actionsFlagSet.Parse(os.Args[1:])

	log.SetLogLevel(*logLevelFlag)

	if *replaceWithFlag != "" {
		fmt.Println("--replace-with", *replaceWithFlag)

		fromVersion := sys_domain.VersionFromString(*replaceWithFlag)
		if fromVersion == nil {
			fmt.Println("Invalid version", *replaceWithFlag)
			os.Exit(1)
		}

		releaseTags, status := sysService.ListReleaseTags()
		if status.Err() {
			fmt.Println("Error checking release tags:", status.Message)
			os.Exit(1)
		}

		for _, tag := range releaseTags {
			if tag.Tag != *replaceWithFlag {
				continue
			}
			sysService.InstallReleaseTag(&tag)
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

	if *verbose {
		for _, key := range os.Environ() {
			fmt.Println(key)
		}
	}
}
