package sys_domain

import "time"

type SystemStatus struct {
	OnlineSince     *time.Time `json:"online_since"`
	ProgramVersion  Version    `json:"version"`
	DatabaseVersion Version    `json:"database_version"`
	FreeSpace       uint64     `json:"free_space"`
	TotalSpace      uint64     `json:"total_space"`
}
