// exempt from testing
package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) GetPackages() (map[string]sys_domain.Package, *core.Status) {
	packages, status := s.repository.GetPackages()
	if status.Err() {
		return nil, &status
	}

	mappedPackages := make(map[string]sys_domain.Package)
	for i := range packages {

		mappedPackages[packages[i].Name] = packages[i]
	}

	return mappedPackages, &status
}
