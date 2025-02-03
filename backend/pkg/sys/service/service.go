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
	InstallReleaseTag(tag *sys_domain.ReleaseTag) *core.Status
}

type SysService struct {
	Repository sys_repository.SysRepositoryInterface
}

func (s *SysService) MigrateDatabase(toVersion *sys_domain.Version) *core.Status {
	status := s.Repository.MigrateDatabase(toVersion)

	return &status
}

func NewSysService(repo sys_repository.SysRepositoryInterface) SysServiceInterface {
	return &SysService{
		Repository: repo,
	}
}
