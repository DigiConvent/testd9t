package sys_repository

import (
	"fmt"
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
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
	err = os.Setenv(constants.TELEGRAM_BOT_TOKEN, botId)
	if err != nil {
		r.db.Exec("update config set telegram_bot_token = ?", "")
		return *core.InternalError("Failed to set telegram bot token as env var, trying to revert")
	}
	return *core.StatusSuccess()
}
