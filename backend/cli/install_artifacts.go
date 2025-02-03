package cli

import (
	"fmt"
	"os"

	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func InstallArtifacts(ofVersion string, sysService sys_service.SysServiceInterface) {

	fromVersion := sys_domain.VersionFromString(ofVersion)
	if fromVersion == nil {
		fmt.Println("Invalid version", ofVersion)
		os.Exit(1)
	}

	releaseTags, status := sysService.ListReleaseTags()
	if status.Err() {
		fmt.Println("Error checking release tags:", status.Message)
		os.Exit(1)
	}

	for _, tag := range releaseTags {
		if tag.Tag != ofVersion {
			continue
		}
		sysService.InstallArtifacts(&tag)
	}
}
