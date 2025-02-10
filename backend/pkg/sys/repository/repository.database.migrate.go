package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) MigrateDatabase(toVersion *sys_domain.Version) core.Status {
	var status core.Status
	if toVersion == nil {
		toVersion, _, status = r.GetCurrentVersion()
		if status.Err() {
			return *core.InternalError(status.Message)
		}
	}

	packages := db.ListPackages()

	for _, pkg := range packages {
		status := r.MigratePackage(pkg, toVersion)
		if status.Err() && status.Code != 404 {
			return status
		}
	}

	_, err := r.db.Exec("insert into versions (major, minor, patch) values (?, ?, ?)", toVersion.Major, toVersion.Minor, toVersion.Patch)
	if err != nil {
		return *core.InternalError("Could not register database version " + toVersion.String())
	}

	return *core.StatusSuccess()
}
