package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetConfiguration() (*sys_domain.Configuration, core.Status) {
	row := r.db.QueryRow("select telegram_bot_token, domain, maintenance from configurations limit 1")

	config := sys_domain.Configuration{}

	err := row.Scan(&config.TelegramBotToken, &config.Domain, &config.Maintenance)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	sys_domain.SetMaintenance(config.Maintenance)

	return &config, *core.StatusSuccess()
}
