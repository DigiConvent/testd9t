package sys_service

import (
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
)

func (s *SysService) SetBotToken(botId string) *core.Status {
	status := s.repository.SetBotToken(botId)
	if !status.Err() {
		os.Setenv(constants.TELEGRAM_BOT_TOKEN, botId)
	}
	return &status
}
