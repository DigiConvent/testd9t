package cli

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	cli_helpers "github.com/DigiConvent/testd9t/cli/helpers"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func HandleFlags() {
	sysService := sys_service.NewSysService(sys_repository.NewSysRepository(db.NewTestSqliteDB("sys")))

	actionsFlagSet := flag.NewFlagSet("Options", flag.ExitOnError)
	verbose := actionsFlagSet.Bool("verbose", false, "Run more verbosely")
	forceFlag := actionsFlagSet.Bool("force", false, "Apply fixes upon a failure during the installation")
	helpFlag := actionsFlagSet.Bool("help", false, "Prints this help message")
	installFlag := actionsFlagSet.String("install", "", "Install ")
	installUsingPresetsFlag := actionsFlagSet.Bool("install-with-presets", false, "Install using presets from a previous installation (usually stored under .d9t-presets)")
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

		releaseTags, status := sysService.ListReleaseTags()
		if status.Err() {
			fmt.Println("Error checking release tags:", status.Message)
			os.Exit(1)
		}

		for _, tag := range releaseTags {
			if tag.Tag != *replaceWithFlag {
				continue
			}
			sysService.InstallArtifacts(&tag)
		}
	}

	if *resetDBFlag {
		ResetDB()
	}

	if *helpFlag || actionsFlagSet.NFlag() == 0 {
		actionsFlagSet.Usage()
		os.Exit(0)
	}

	if *versionsFlag {
		cli_helpers.ListVersions(sysService)
	}

	if *listFlavoursFlag != "" {
		ListFlavours(sysService)
	}

	if *statusFlag {
		ShowStatus(sysService)
	}

	if *installFlag != "" {
		status := InstallArtifacts(sys_domain.ProgramVersion, sysService)
		if status.Err() {
			fmt.Println("Error installing artifacts:", status.Message)
		}
		Install(sysService, installFlag, *forceFlag, *verbose, *installUsingPresetsFlag)
		err := os.RemoveAll("/tmp/testd9t/")
		if err != nil {
			fmt.Println(err)
		}

		err = exec.Command("chown", "-R", "testd9t:testd9t", constants.HOME_PATH).Run()
		if err != nil {
			log.Error("Error setting ownership for artifacts: " + err.Error())
		} else {
			log.Success("Set ownership for artifacts")
		}
	}

	if *migrateDBFlag {
		MigrateDB(sysService)
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
