package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/DigiConvent/testd9t/install"
	"github.com/DigiConvent/testd9t/install/version"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if version.Dev() {
		godotenv.Load(".env")
	} else {
		godotenv.Load("/etc/digiconvent/env")
	}

	handleFlags()

	router := gin.Default()

	router.NoRoute(handleFrontend())
	router.Run(":8080")
}

func handleFrontend() gin.HandlerFunc {
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
}

func handleFlags() {
	actionsFlagSet := flag.NewFlagSet("Options", flag.ExitOnError)
	createBackupFlag := actionsFlagSet.String("backup", "", "Creates a backup of the database and uploaded files")
	envFlag := actionsFlagSet.String("env", "", "Path to the env file to overwrite the default one")
	forceFlag := actionsFlagSet.Bool("force", false, "Pseudo sudo")
	helpFlag := actionsFlagSet.Bool("help", false, "Prints this help message")
	installFlag := actionsFlagSet.Bool("install", false, "Installs the program")
	resetDBFlag := actionsFlagSet.Bool("reset-db", false, "Resets the database")
	restoreBackupFlag := actionsFlagSet.String("restore", "", "Restores the database and uploaded files from a backup")
	runFlag := actionsFlagSet.Bool("run", false, "Deploy")
	verboseFlag := actionsFlagSet.Bool("verbose", false, "Prints more information")
	versionFlag := actionsFlagSet.Bool("version", false, "Prints the current version")
	versionsFlag := actionsFlagSet.Bool("versions", false, "List all available versions")

	actionsFlagSet.Parse(os.Args[1:])

	if *helpFlag {
		actionsFlagSet.Usage()
		os.Exit(0)
	}

	if *envFlag != "" {
		install.OverwriteFromFile(*envFlag, *verboseFlag)
	}

	if *createBackupFlag != "" {
		install.CreateBackup(createBackupFlag)
	}

	if *installFlag {
		fmt.Println("Using force:", *forceFlag)
		install.Install(*forceFlag)
	}

	if *restoreBackupFlag != "" {
		install.RestoreBackup(*restoreBackupFlag)
	}

	if *resetDBFlag {
		install.ResetDB()
		install.MigrateTo(*version.VersionFromString(version.ProgramVersion()))
	}

	if *versionsFlag {
		fmt.Println("\n--versions")
		versions, err := version.ListReleases()
		if err != nil {
			fmt.Println("Error fetching versions:", err)
			os.Exit(1)
		}
		for _, tag := range versions {
			add := ""
			if *tag.TagName == version.ProgramVersion() {
				add += "program"
			}
			if *tag.TagName == version.MigrationVersion().String() {
				if add != "" {
					add += "/"
				}
				add += "migration"
			}
			if add != "" {
				add = " (current " + add + " version)"
			}
			fmt.Println(*tag.TagName, add)
		}
	}

	if *versionFlag {
		fmt.Println("\n--version")
		if version.Dev() {
			fmt.Println("Development")
		} else {
			fmt.Println("Production")
		}

		currentVersion := version.MigrationVersion()
		fmt.Println("Migration  :", currentVersion.String())
		fmt.Println("Environment:", os.Getenv("VERSION"))
	}

	if !*runFlag {
		os.Exit(0)
	}
}
