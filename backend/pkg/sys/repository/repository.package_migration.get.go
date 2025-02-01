package sys_repository

import (
	"os"
	"path"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/file_repo"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (r *SysRepository) GetPackageMigrationScript(pkgName string, getVersion sys_domain.Version) (string, core.Status) {
	var script string
	if sys_domain.ProgramVersion == "dev" {
		projectPath := sys_domain.DevPath()
		firstVersionPath := path.Join(projectPath, "pkg", pkgName, "db", getVersion.String())
		files, err := os.ReadDir(firstVersionPath)
		if err != nil {
			return "", *core.InternalError(err.Error())
		}

		for _, file := range files {
			migrationPath := path.Join(firstVersionPath, file.Name())
			contents, err := os.ReadFile(migrationPath)
			if err != nil {
				return "", *core.InternalError(err.Error())
			}

			script += string(contents)
		}
	} else {
		contents, err := file_repo.NewRepoRemote().ReadRawFile(".meta/migrations/" + getVersion.String() + "/" + pkgName + ".sql")
		if err != nil {
			if strings.HasSuffix(err.Error(), "404 Not Found") {
				return "", *core.NotFoundError("")
			}
			return "", *core.InternalError(err.Error())
		}

		script = string(contents)
	}

	return script, *core.StatusSuccess()
}
