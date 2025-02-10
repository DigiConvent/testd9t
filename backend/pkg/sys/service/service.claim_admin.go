// exempt from testing
package sys_service

import "github.com/DigiConvent/testd9t/core"

func (s *SysService) ClaimAdmin(telegramId string) *core.Status {
	status := s.repository.ClaimAdmin(telegramId)
	return &status
}
