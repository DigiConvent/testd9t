package sys_service

import (
	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
)

type SysServiceInterface interface {
	Init() *core.Status
	GetSystemStatus() (*sys_domain.SystemStatus, *core.Status)
	ClaimAdmin(telegramId string) *core.Status

	GetConfiguration() (*sys_domain.Configuration, *core.Status)
	GetPackages() (map[string]sys_domain.Package, *core.Status)
	GetPackageVersions(pkgName string) ([]sys_domain.Version, *core.Status)

	MigrateDatabase(toVersion *sys_domain.Version) *core.Status
	MigratePackage(pkgName string, toVersion *sys_domain.Version) *core.Status

	ListFlavours() ([]string, *core.Status)

	ListReleaseTags() ([]sys_domain.ReleaseTag, *core.Status)
	InstallArtifacts(tag *sys_domain.ReleaseTag) *core.Status
}

type SysService struct {
	repository sys_repository.SysRepositoryInterface
}

func (s *SysService) MigrateDatabase(toVersion *sys_domain.Version) *core.Status {
	status := s.repository.MigrateDatabase(toVersion)

	return &status
}

func NewSysService(repo sys_repository.SysRepositoryInterface) SysServiceInterface {
	return &SysService{
		repository: repo,
	}
}
