package sys_service

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
)

func (s *SysService) Init() *core.Status {
	if s.Repository.IsInitialised() {
		return core.ConflictError("Already initialised")
	}

	initStatus := s.Repository.InitDatabase()

	if initStatus.Err() {
		return core.InternalError(initStatus.Message)
	}

	sysStatus, status := s.GetSystemStatus()

	if status.Err() {
		return core.InternalError(status.Message)
	}

	packages := db.ListPackages()

	for _, pkg := range packages {
		status := s.MigratePackage(pkg, sysStatus.ProgramVersion)
		if status.Err() && status.Code != 404 {
			fmt.Println("Error migrating package", pkg, ":", status.Message)
		}
	}

	return core.StatusSuccess()
}
