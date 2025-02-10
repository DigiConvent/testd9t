// exempt from testing
package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) MigratePackage(pkgName string, toVersion *sys_domain.Version) *core.Status {
	status := s.repository.MigratePackage(pkgName, toVersion)
	if status.Err() {
		return &status
	}
	return core.StatusSuccess()
}
