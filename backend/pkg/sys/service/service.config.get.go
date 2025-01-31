package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) GetConfiguration() (*sys_domain.Configuration, *core.Status) {
	config, status := s.Repository.GetConfiguration()
	if status.Err() {
		return nil, &status
	}
	return config, &status
}
