package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) IsInitialised() bool {
	resultSet := r.DB.QueryRow("SELECT major, minor, patch FROM packages WHERE name = 'sys'")
	var major, minor, patch int = -1, -1, -1
	err := resultSet.Scan(&major, &minor, &patch)
	return err == nil && major >= 0 && minor >= 0 && patch >= 0
}

func (r *SysRepository) InitDatabase() core.Status {
	var script string

	script, status := r.GetPackageMigrationScript("sys", sys_domain.Version{Major: 0, Minor: 0, Patch: 0})

	if status.Err() {
		return status
	}

	_, err := r.DB.Exec(string(script))
	if err != nil {
		return *core.InternalError("Could not initialise")
	}

	_, err = r.DB.Exec("INSERT INTO packages (name, major, minor, patch) VALUES ('sys', 0, 0, 0)")

	if err != nil {
		return *core.InternalError("Could not register initialisation")
	}

	return *core.StatusSuccess()
}
