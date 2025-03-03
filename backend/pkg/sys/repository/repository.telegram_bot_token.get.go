package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
)

func (r *SysRepository) GetBotToken() (string, core.Status) {
	var token string
	err := r.db.QueryRow("select telegram_bot_token from configurations limit 1").Scan(&token)
	if err != nil {
		return "", *core.InternalError(err.Error())
	}
	return token, *core.StatusSuccess()
}
