package cli_helpers

import (
	"fmt"
	"os"

	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func ListVersions(sysService sys_service.SysServiceInterface) {
	sysStatus, status := sysService.GetSystemStatus()

	if status.Err() {
		fmt.Println("Error fetching system status:", status.Message)
		os.Exit(1)
	}

	versions, status := sysService.ListReleaseTags()
	if status.Err() {
		fmt.Println("Error fetching versions:", status.Message)
		os.Exit(1)
	}

	if status.Err() {
		fmt.Println("Error fetching system status:", status.Message)
		os.Exit(1)
	}

	if sysStatus.Version.ProgramVersion.Major == -1 {
		fmt.Println("Program version: dev")
		if sysStatus.Version.DatabaseVersion.Equals(&sysStatus.Version.ProgramVersion) {
			fmt.Println("No migrations found, run with --migrate-db to migrate to a compatible version")
		}
	}

	fmt.Println("┏━━━━━━━━━┳━━━━━━━━━┳━━━━━━━━━━━┓")
	fmt.Println("┃ Version ┃ Program ┃ Migration ┃")

	for _, tag := range versions {
		var program, migrationExecuted string
		if tag.Tag == sysStatus.Version.ProgramVersion.String() {
			program = "✓"
		}
		if tag.Tag == sysStatus.Version.DatabaseVersion.String() {
			migrationExecuted = "✓"
		}
		fmt.Println("┣━━━━━━━━━╋━━━━━━━━━╋━━━━━━━━━━━┫")
		fmt.Printf("┃ %7s ┃ %7s ┃ %9s ┃\n", tag.Tag, program, migrationExecuted)
	}
	fmt.Println("┗━━━━━━━━━┻━━━━━━━━━┻━━━━━━━━━━━┛")
}
