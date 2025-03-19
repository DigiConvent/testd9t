package sys_service

import (
	"os"
	"os/exec"

	"github.com/DigiConvent/testd9t/core"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
)

type SysServiceInterface interface {
	Init() *core.Status
	RefreshInstallation() *core.Status
	GetSystemStatus() (*sys_domain.SystemStatus, *core.Status)

	GetConfiguration() (*sys_domain.Configuration, *core.Status)
	GetPackages() (map[string]sys_domain.Package, *core.Status)
	GetPackageVersions(pkgName string) ([]sys_domain.Version, *core.Status)

	MigratePackageDatabases(toVersion *sys_domain.Version) *core.Status
	MigratePackage(pkgName string, toVersion *sys_domain.Version) *core.Status

	ListFlavours() ([]string, *core.Status)

	ListReleaseTags() ([]sys_domain.ReleaseTag, *core.Status)
	InstallArtifacts(tag *sys_domain.ReleaseTag) *core.Status

	SetBotToken(botId string) *core.Status
	SetDomain(domain string) *core.Status
	SetLargeLogo(data []byte) *core.Status
	SetSmallLogo(data []byte) *core.Status
}

type SysService struct {
	repository sys_repository.SysRepositoryInterface
}

func (s *SysService) RefreshInstallation() *core.Status {
	currentVersion := sys_domain.ProgramVersion
	flavour, err := s.repository.GetFlavour()
	if err != nil {
		return core.InternalError(err.Error())
	}

	supportedFlavours, status := s.repository.ListFlavoursForVersion()

	if status.Err() {
		return &status
	}

	if !core_utils.Contains(supportedFlavours, flavour) {
		return core.InternalError("Flavour " + flavour + " is not supported for version " + currentVersion)
	}

	if currentVersion == "0.0.0" {
		cmds := []string{
			"rm -rf /home/testd9t",
			"wget https://github.com/DigiConvent/testd9t/releases/download/0.0.0/main -o /tmp/main",
			"chmod +x /tmp/main",
			"./tmp/main " + flavour + " --install --presets=/tmp/.d9t_presets",
		}

		for _, cmd := range cmds {
			err = exec.Command(cmd).Run()
			if err != nil {
				return core.InternalError(err.Error())
			}
		}
		os.Exit(0)
	} else {
		// TODO some later time
	}

	return core.StatusSuccess()
}

func (s *SysService) MigratePackageDatabases(toVersion *sys_domain.Version) *core.Status {
	status := s.repository.MigrateDatabase(toVersion)
	return &status
}

func NewSysService(repo sys_repository.SysRepositoryInterface) SysServiceInterface {
	return &SysService{
		repository: repo,
	}
}
