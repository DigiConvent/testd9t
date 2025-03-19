package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
)

func (r *SysRepository) SetBotToken(botId string) core.Status {
	if botId == "" {
		return *core.BadRequestError("Bot token cannot be empty")
	}

	_, err := r.db.Exec("update configurations set telegram_bot_token = ?", botId)
	if err != nil {
		log.Error(err.Error())
		return *core.InternalError("Failed to set telegram bot token")
	}
	return *core.StatusSuccess()
}
