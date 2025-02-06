// exempt from testing
package sys_service

import (
	"os/exec"
	"strconv"
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
	var testd9tStat syscall.Statfs_t

	syscall.Statfs("/", &stat)
	syscall.Statfs("/home/testd9t/", &testd9tStat)

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)

	cmd := exec.Command("du", "-sb", "/home/testd9t/")
	out, _ := cmd.Output()

	outString := string(out)

	outBytes, _ := strconv.Atoi(outString)

	return &sys_domain.SystemStatus{
		TotalSpace:      total,
		FreeSpace:       free,
		DataSpace:       uint64(outBytes),
		ProgramVersion:  *programVersion,
		DatabaseVersion: *databaseVersion,
		BuiltAt:         sys_domain.CompiledAt,
		OnlineSince:     sys_domain.StartTime.Format(core_utils.FormattedTime),
	}, core.StatusSuccess()
}
