package sys_service

import "github.com/DigiConvent/testd9t/core"

func (s *SysService) SetBotToken(botId string) *core.Status {
	status := s.repository.SetBotToken(botId)
	return &status
}
