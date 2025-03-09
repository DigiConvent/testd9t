package sys_domain

import "time"

var StartTime time.Time

type SystemStatus struct {
	Dns struct {
		DkimIs      string `json:"dkim_is"`
		DkimShould  string `json:"dkim_should"`
		DmarcIs     string `json:"dmarc_is"`
		DmarcShould string `json:"dmarc_should"`
		DnsIs       string `json:"dns_is"`
		DnsShould   string `json:"dns_should"`
		MxIs        string `json:"mx_is"`
		MxShould    string `json:"mx_should"`
		SpfIs       string `json:"spf_is"`
		SpfShould   string `json:"spf_should"`
	} `json:"dns"`
	Server struct {
		DataSpace  uint64 `json:"data_space"`
		FreeSpace  uint64 `json:"free_space"`
		TotalSpace uint64 `json:"total_space"`
	} `json:"server"`
	TelegramBot struct {
		TelegramBotHint   string `json:"telegram_bot_hint"`
		TelegramBotStatus string `json:"telegram_bot"`
	} `json:"telegram_bot"`
	Version struct {
		BuiltAt         string  `json:"built_at"`
		OnlineSince     string  `json:"online_since"`
		DatabaseVersion Version `json:"database_version"`
		ProgramVersion  Version `json:"version"`
	} `json:"version"`
}
