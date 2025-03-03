package sys_service

import (
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
)

func (s *SysService) SetDomain(domain string) *core.Status {
	status := s.repository.SetDomain(domain)
	if !status.Err() {
		os.Setenv(constants.DOMAIN, domain)
	}
	return &status
}
