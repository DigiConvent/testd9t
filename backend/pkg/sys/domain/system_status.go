package sys_domain

import "time"

var StartTime time.Time

type SystemStatus struct {
	Dns struct {
		DkimHint    string `json:"dkim_hint"`
		DkimStatus  string `json:"dkim_status"`
		DmarcHint   string `json:"dmarc_hint"`
		DmarcStatus string `json:"dmarc_status"`
		DnsHint     string `json:"dns_hint"`
		DnsStatus   string `json:"dns_status"`
		MxHint      string `json:"mx_hint"`
		MxStatus    string `json:"mx_status"`
		SpfHint     string `json:"spf_hint"`
		SpfStatus   string `json:"spf_status"`
	} `json:"dns"`
	Server struct {
		DataSpace  uint64 `json:"data_space"`
		FreeSpace  uint64 `json:"free_space"`
		TotalSpace uint64 `json:"total_space"`
	} `json:"server"`
	TelegramBot struct {
		TelegramBotHint   string `json:"telegram_bot_hint"`
		TelegramBotStatus string `json:"telegram_bot_status"`
	} `json:"telegram_bot"`
	Version struct {
		BuiltAt         string  `json:"built_at"`
		OnlineSince     string  `json:"online_since"`
		DatabaseVersion Version `json:"database_version"`
		ProgramVersion  Version `json:"version"`
	} `json:"version"`
}
