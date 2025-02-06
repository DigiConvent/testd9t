// exempt from testing
package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) GetPackageVersions(pkgName string) ([]sys_domain.Version, *core.Status) {
	versions, status := s.Repository.GetPackageVersions(pkgName)

	if status.Err() {
		return nil, &status
	}

	return versions, &status
}
