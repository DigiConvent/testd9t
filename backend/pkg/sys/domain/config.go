package sys_domain

var config Configuration = Configuration{
	TelegramBotToken: "",
	Domain:           "",
	Maintenance:      false,
}

type Configuration struct {
	TelegramBotToken string `json:"telegram_bot_token"`
	Domain           string `json:"domain"`
	Maintenance      bool   `json:"maintenance"`
}

func IsMaintenance() bool {
	return config.Maintenance
}

func SetMaintenance(maintenance bool) {
	config.Maintenance = maintenance
}
