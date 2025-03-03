package sys_domain

import "time"

var StartTime time.Time

type SystemStatus struct {
	OnlineSince     string  `json:"online_since"`
	ProgramVersion  Version `json:"version"`
	DatabaseVersion Version `json:"database_version"`
	FreeSpace       uint64  `json:"free_space"`
	TotalSpace      uint64  `json:"total_space"`
	DataSpace       uint64  `json:"data_space"`
	BuiltAt         string  `json:"built_at"`

	// telegram mini app stuff
	TelegramBotStatus string `json:"telegram_bot_status"`
	TelegramBotHint   string `json:"telegram_bot_hint"`

	// dns
	DnsStatus   string `json:"dns_status"`
	DnsHint     string `json:"dns_hint"`
	MxStatus    string `json:"mx_status"`
	MxHint      string `json:"mx_hint"`
	DmarcStatus string `json:"dmarc_status"`
	DmarcHint   string `json:"dmarc_hint"`
	DkimStatus  string `json:"dkim_status"`
	DkimHint    string `json:"dkim_hint"`
	SpfStatus   string `json:"spf_status"`
	SpfHint     string `json:"spf_hint"`
}
