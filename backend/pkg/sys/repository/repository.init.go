package sys_repository

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
)

func (r *SysRepository) SetBotToken(botId string) core.Status {
	if botId == "" {
		return *core.BadRequestError("Bot token cannot be empty")
	}

	_, err := r.db.Exec("update config set telegram_bot_token = ?", botId)
	if err != nil {
		fmt.Println(err.Error())
		return *core.InternalError("Failed to set telegram bot token")
	}
	return *core.StatusSuccess()
}
