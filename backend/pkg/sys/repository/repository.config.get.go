package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetConfiguration() (*sys_domain.Configuration, core.Status) {
	row := r.DB.QueryRow("SELECT telegram_bot_token, domain FROM configuration LIMIT 1")

	config := sys_domain.Configuration{}

	err := row.Scan(&config.TelegramBotToken, &config.Domain)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	return &config, *core.StatusSuccess()
}
