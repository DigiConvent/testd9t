package sys_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

type SysRepositoryInterface interface {
	GetCurrentVersion() (*sys_domain.Version, *sys_domain.Version, core.Status)
	ListFlavoursForVersion() ([]string, core.Status)

	SetBotToken(botId string) core.Status
	GetBotToken() (string, core.Status)

	SetDomain(domain string) core.Status
	GetConfiguration() (*sys_domain.Configuration, core.Status)

	IsInitialised() bool
	InitDatabase() core.Status
	MigrateDatabase(toVersion *sys_domain.Version) core.Status

	GetPackages() ([]sys_domain.Package, core.Status)
	GetPackageVersions(pkgName string) ([]sys_domain.Version, core.Status)
	GetPackageMigrationScript(pkgName string, fromVersion *sys_domain.Version) (string, core.Status)
	MigratePackage(pkgName string, toVersion *sys_domain.Version) core.Status

	ListReleaseTags() ([]sys_domain.ReleaseTag, *core.Status)

	GetDiskUsage() (*sys_domain.DiskUsage, *core.Status)
}

type SysRepository struct {
	db db.DatabaseInterface
}

func NewSysRepository(conn db.DatabaseInterface) SysRepositoryInterface {
	return &SysRepository{
		db: conn,
	}
}
