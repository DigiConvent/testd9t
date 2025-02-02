package sys_domain

import "time"

var StartTime time.Time

type SystemStatus struct {
	OnlineSince     string  `json:"online_since"`
	ProgramVersion  Version `json:"version"`
	DatabaseVersion Version `json:"database_version"`
	FreeSpace       uint64  `json:"free_space"`
	TotalSpace      uint64  `json:"total_space"`
	BuiltAt         string  `json:"built_at"`
}
