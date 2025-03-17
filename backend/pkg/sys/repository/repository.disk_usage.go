package sys_repository

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func getDirSize(dir string) int64 {
	var size int64
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", path, err)
			return filepath.SkipDir
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err == nil {
				size += info.Size()
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", dir, err)
	}
	return size
}

func (r *SysRepository) GetDiskUsage() (*sys_domain.DiskUsage, *core.Status) {
	result := sys_domain.DiskUsage{}

	dbPath := constants.HOME_PATH + "db/"

	result.Data.IamSize = int(getDirSize(dbPath + "iam"))
	result.Data.SysSize = int(getDirSize(dbPath + "sys"))
	result.Data.PostSize = int(getDirSize(dbPath + "post"))
	result.Data.Certificates = int(getDirSize(constants.HOME_PATH + "certs"))
	result.Program.Backend = int(getDirSize(constants.HOME_PATH + "backend"))
	result.Program.Frontend = int(getDirSize(constants.HOME_PATH + "frontend"))

	var totalSpace uint64
	var freeSpace uint64
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)

	if err != nil {
		totalSpace = 0
		freeSpace = 0
	} else {
		totalSpace = stat.Blocks * uint64(stat.Bsize)
		freeSpace = stat.Bfree * uint64(stat.Bsize)
	}

	result.OS = int(totalSpace-freeSpace) - result.Data.IamSize - result.Data.SysSize - result.Data.PostSize
	result.TotalServer = int(totalSpace)
	result.TotalHome = int(getDirSize(constants.HOME_PATH))
	result.Free = int(freeSpace)

	return &result, core.StatusSuccess()
}
