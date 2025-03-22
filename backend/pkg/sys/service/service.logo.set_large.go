// exempt from testing

package sys_service

import (
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) SetLargeLogo(data []byte) *core.Status {
	return setLogo("large", data)
}

func setLogo(variant string, data []byte) *core.Status {
	if variant != "large" && variant != "small" {
		return core.BadRequestError("Invalid variant")
	}

	logoFolder := constants.HOME_PATH + "data/sys/logo/"
	logoPath := logoFolder + variant + ".jpg"
	symlinkPath := constants.HOME_PATH + "frontend/assets/" + variant + ".jpg"
	os.MkdirAll(logoFolder, 0755)

	if sys_domain.ProgramVersion != "dev" {
		info, _ := os.Lstat(symlinkPath)
		if info != nil {
			os.Remove(symlinkPath)
		}
	}

	os.Remove(logoPath)
	err := os.WriteFile(logoPath, data, 0644)
	if err != nil {
		return core.InternalError("Could not write logo:" + err.Error())
	}

	if sys_domain.ProgramVersion != "dev" {
		if err := os.Symlink(logoPath, symlinkPath); err != nil {
			return core.InternalError(err.Error())
		}
	}

	return core.StatusSuccess()
}
