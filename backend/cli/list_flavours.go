package cli

import (
	"fmt"
	"os"

	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func ListFlavours(sysService sys_service.SysServiceInterface) {
	fmt.Println("--supported-flavours")
	flavours, status := sysService.ListFlavours()
	if status.Err() {
		fmt.Println("Error fetching flavours:", status.Message)
		os.Exit(1)
	}

	if len(flavours) == 0 {
		fmt.Println("No flavours found")
		os.Exit(0)
	}

	for _, flavour := range flavours {
		fmt.Println(flavour)
	}
}
