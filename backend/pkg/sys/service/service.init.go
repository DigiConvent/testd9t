// exempt from testing
package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (s *SysService) Init() *core.Status {
	if s.repository.IsInitialised() {
		return core.ConflictError("Already initialised")
	}

	initStatus := s.repository.InitDatabase()

	if initStatus.Err() {
		return core.InternalError(initStatus.Message)
	}

	return core.StatusSuccess()
}
