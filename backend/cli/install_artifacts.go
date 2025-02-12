package cli

import (
	"os"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func InstallArtifacts(ofVersion string, sysService sys_service.SysServiceInterface) *core.Status {
	fromVersion := sys_domain.VersionFromString(ofVersion)
	if fromVersion == nil {
		log.Error("Invalid version" + ofVersion)
		os.Exit(1)
	}

	releaseTags, status := sysService.ListReleaseTags()
	if status.Err() {
		log.Error("Error checking release tags: " + status.Message)
		os.Exit(1)
	}

	var tag sys_domain.ReleaseTag
	for _, existingTag := range releaseTags {
		if existingTag.Tag != ofVersion {
			continue
		} else {
			tag = existingTag
		}
	}
	return sysService.InstallArtifacts(&tag)
}
