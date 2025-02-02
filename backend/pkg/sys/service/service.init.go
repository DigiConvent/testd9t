package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (s *SysService) Init() *core.Status {
	if s.Repository.IsInitialised() {
		return core.ConflictError("Already initialised")
	}

	initStatus := s.Repository.InitDatabase()

	if initStatus.Err() {
		return core.InternalError(initStatus.Message)
	}

	return core.StatusSuccess()
}
