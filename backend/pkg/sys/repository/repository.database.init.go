package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) IsInitialised() bool {
	resultSet := r.db.QueryRow("select count(*) from sqlite_master where type='table' and name=?", "versions")
	var count = 0
	err := resultSet.Scan(&count)
	if err != nil || count == 0 {
		return false
	}
	return true
}

func (r *SysRepository) InitDatabase() core.Status {
	var script string

	script, status := r.GetPackageMigrationScript("sys", &sys_domain.Version{Major: 0, Minor: 0, Patch: 0})

	if status.Err() {
		return status
	}

	_, err := r.db.Exec(string(script))
	if err != nil {
		return *core.InternalError("Could not initialise " + err.Error())
	}

	_, err = r.db.Exec("insert into packages (name, major, minor, patch) values ('sys', 0, 0, 0)")

	if err != nil {
		return *core.InternalError("Could not register sys:0.0.0 " + err.Error())
	} else {
		log.Success("Registered sys:0.0.0")
	}

	return *core.StatusSuccess()
}
