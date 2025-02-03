package cli

import (
	"fmt"

	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func MigrateDB(sysService sys_service.SysServiceInterface) {
	fmt.Println("--migrate-db")
	sysService.MigrateDatabase(nil)
}
