package sys_repository

import (
	"strconv"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) IsInitialised() bool {
	resultSet := r.db.QueryRow("SELECT major, minor, patch FROM packages WHERE name = 'sys'")
	var major, minor, patch int = -1, -1, -1
	err := resultSet.Scan(&major, &minor, &patch)
	if err != nil {
		return false
	} else {
		log.Success("System is initialised at version " + strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch))
	}
	return err == nil && major >= 0 && minor >= 0 && patch >= 0
}

func (r *SysRepository) InitDatabase() core.Status {
	var script string

	script, status := r.GetPackageMigrationScript("sys", &sys_domain.Version{Major: 0, Minor: 0, Patch: 0})

	if status.Err() {
		return status
	}

	_, err := r.db.Exec(string(script))
	if err != nil {
		return *core.InternalError("Could not initialise")
	}

	_, err = r.db.Exec("INSERT INTO packages (name, major, minor, patch) VALUES ('sys', 0, 0, 0)")

	if err != nil {
		return *core.InternalError("Could not register sys:0.0.0 " + err.Error())
	} else {
		log.Success("Registered sys:0.0.0")
	}

	return *core.StatusSuccess()
}
