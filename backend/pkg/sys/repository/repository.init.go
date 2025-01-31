package sys_repository

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
)

func (r *SysRepository) SetBotToken(botId string) core.Status {
	if botId == "" {
		return *core.BadRequestError("Bot token cannot be empty")
	}

	// status := r.UpdateToVersion(&sys_domain.Version{Major: 0, Minor: 0, Patch: 0})
	// if status.Err() {
	// 	return *core.InternalError("Failed to migrate to version 0.0.0" + status.Message)
	// }

	_, err := r.DB.Exec("update config set telegram_bot_token = $1", botId)
	if err != nil {
		fmt.Println(err.Error())
		return *core.InternalError("Failed to set telegram bot token")
	}
	return *core.StatusSuccess()
}
