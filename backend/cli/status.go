package cli

import (
	"fmt"
	"os"

	cli_helpers "github.com/DigiConvent/testd9t/cli/helpers"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func ShowStatus(sysService sys_service.SysServiceInterface) {
	sysStatus, status := sysService.GetSystemStatus()
	if status.Err() {
		fmt.Println("Error getting system status:", status.Message)
		os.Exit(1)
	}
	fmt.Println("--status")
	if sys_domain.ProgramVersion == "dev" {
		fmt.Println("Development")
	} else {
		fmt.Println("Production")
	}

	currentVersion := &sysStatus.Version.DatabaseVersion

	fmt.Println("Migration  :", currentVersion.String())

	fmt.Println("Environment:", sysStatus.Version.ProgramVersion.String())

	cli_helpers.ListPackages(sysService)
}
