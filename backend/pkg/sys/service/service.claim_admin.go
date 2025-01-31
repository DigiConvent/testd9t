package sys_service

import "github.com/DigiConvent/testd9t/core"

func (s *SysService) ClaimAdmin(telegramId string) *core.Status {
	status := s.Repository.ClaimAdmin(telegramId)
	return &status
}
