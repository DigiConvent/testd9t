package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetCurrentVersion() (*sys_domain.Version, *sys_domain.Version, core.Status) {
	var programVersion *sys_domain.Version = sys_domain.VersionFromString(sys_domain.ProgramVersion)
	if programVersion == nil {
		programVersion = &sys_domain.Version{Major: -1, Minor: -1, Patch: -1}
	}

	row := r.DB.QueryRow(`SELECT major, minor, patch FROM versions ORDER BY major DESC, minor DESC, patch DESC LIMIT 1`)

	var result = &sys_domain.Version{
		Major: -1,
		Minor: -1,
		Patch: -1,
	}

	err := row.Scan(&result.Major, &result.Minor, &result.Patch)

	if err != nil {
		return programVersion, result, *core.StatusSuccess()
	}

	return programVersion, result, *core.StatusSuccess()
}
