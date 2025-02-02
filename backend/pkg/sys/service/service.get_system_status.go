package sys_service

import (
	"syscall"

	"github.com/DigiConvent/testd9t/core"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) GetSystemStatus() (*sys_domain.SystemStatus, *core.Status) {
	programVersion, databaseVersion, status := s.Repository.GetCurrentVersion()
	if status.Err() {
		programVersion = &sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
		databaseVersion = &sys_domain.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
	}
	var stat syscall.Statfs_t

	syscall.Statfs("/", &stat)

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)

	return &sys_domain.SystemStatus{
		TotalSpace:      total,
		FreeSpace:       free,
		ProgramVersion:  *programVersion,
		DatabaseVersion: *databaseVersion,
		BuiltAt:         sys_domain.CompiledAt,
		OnlineSince:     sys_domain.StartTime.Format(core_utils.FormattedTime),
	}, core.StatusSuccess()
}
