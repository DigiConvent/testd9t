package sys_service

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
)

func (s *SysService) Init() *core.Status {
	if s.Repository.IsInitialised() {
		return core.StatusSuccess()
	}

	status := s.Repository.InitDatabase()

	if status.Err() {
		return core.InternalError(status.Message)
	}
	fmt.Println("Initialised")
	return &status
}
