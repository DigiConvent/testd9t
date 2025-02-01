package cli

import (
	"fmt"
	"os"

	"github.com/DigiConvent/testd9t/core/db"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func MigrateDB(sysService sys_service.SysServiceInterface) {
	fmt.Println("--migrate-db")
	sysStatus, status := sysService.GetSystemStatus()

	if status.Err() {
		fmt.Println("Error getting system status:", status.Message)
		os.Exit(1)
	}
	packages := db.ListPackages()

	for _, pkg := range packages {
		status := sysService.MigratePackage(pkg, sysStatus.ProgramVersion)
		if status.Err() && status.Code != 404 {
			fmt.Println("Error migrating package", pkg, ":", status.Message)
		}
	}
}
