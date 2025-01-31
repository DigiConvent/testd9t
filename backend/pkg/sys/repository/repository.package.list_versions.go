package sys_repository

import (
	"os"
	"path"

	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetPackageVersions(pkgName string) ([]sys_domain.Version, core.Status) {
	var versions []sys_domain.Version
	if sys_domain.ProgramVersion == "dev" {
		packagePath := path.Join(sys_domain.DevPath(), "pkg", pkgName, "db")

		versionDirs, err := os.ReadDir(packagePath)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		for _, versionDir := range versionDirs {
			if versionDir.IsDir() {
				versions = append(versions, *sys_domain.VersionFromString(versionDir.Name()))
			}
		}
	} else {
		releaseTags, status := r.ListReleaseTags()
		if status.Err() {
			return nil, *status
		}

		for _, releaseTag := range releaseTags {
			for _, migrationName := range releaseTag.Migrations {
				if migrationName == pkgName+".sql" {
					versions = append(versions, *sys_domain.VersionFromString(releaseTag.Tag))
				}
			}
		}
	}

	return versions, *core.StatusSuccess()
}
