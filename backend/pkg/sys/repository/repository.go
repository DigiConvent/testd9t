package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

type SysRepositoryInterface interface {
	ListVersions() ([]sys_domain.Version, core.Status)
	GetCurrentVersion() (*sys_domain.Version, *sys_domain.Version, core.Status)
	ListFlavoursForVersion() ([]string, core.Status)

	SetBotToken(botId string) core.Status
	ClaimAdmin(telegramId string) core.Status

	GetConfiguration() (*sys_domain.Configuration, core.Status)
	InitDatabase() core.Status
	IsInitialised() bool

	GetPackages() ([]sys_domain.Package, core.Status)
	GetPackageVersions(pkgName string) ([]sys_domain.Version, core.Status)
	GetPackageMigrationScript(pkgName string, fromVersion sys_domain.Version) (string, core.Status)
	MigratePackage(pkgName string, toVersion sys_domain.Version) core.Status

	ListReleaseTags() ([]sys_domain.ReleaseTag, *core.Status)
}

type SysRepository struct {
	DB          db.DatabaseInterface
	GithubToken *string
}

func (r *SysRepository) ListVersions() ([]sys_domain.Version, core.Status) {
	return []sys_domain.Version{
		{Major: 0, Minor: 0, Patch: 0},
		{Major: 0, Minor: 0, Patch: 1},
		{Major: 0, Minor: 0, Patch: 2},
		{Major: 0, Minor: 0, Patch: 3},
	}, *core.StatusSuccess()
}
func NewSysRepository(conn db.DatabaseInterface) SysRepositoryInterface {
	return &SysRepository{
		DB: conn,
	}
}
