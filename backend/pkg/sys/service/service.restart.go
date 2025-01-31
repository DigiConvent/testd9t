package sys_service

import "os"

func (s *SysService) Restart() {
	os.Exit(0)
}
