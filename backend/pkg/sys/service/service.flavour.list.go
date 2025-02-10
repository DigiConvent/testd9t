// exempt from testing
package sys_service

import "github.com/DigiConvent/testd9t/core"

func (s *SysService) ListFlavours() ([]string, *core.Status) {
	flavours, status := s.repository.ListFlavoursForVersion()
	if status.Err() {
		return nil, &status
	}

	return flavours, &status
}
