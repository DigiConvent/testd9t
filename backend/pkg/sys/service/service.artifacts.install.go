// exempt from testing
package sys_service

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

const targetBinaryPath = constants.HOME_PATH + "backend/main"
const targetFrontendPath = constants.HOME_PATH + "frontend"

const sourceFrontendZipPath = "/tmp/testd9t/frontend.zip"

func (s *SysService) InstallArtifacts(tag *sys_domain.ReleaseTag) *core.Status {
	log.Info("Installing artifacts for version " + tag.Tag)
	var err error

	err = tag.DownloadAsset("main", targetBinaryPath)
	if err != nil {
		return core.InternalError("Error downloading binary for version: " + tag.Tag + " " + err.Error())
	}

	err = exec.Command("chmod", "+x", targetBinaryPath).Run()
	if err != nil {
		return core.InternalError("Error setting permissions for binary: " + err.Error())
	}

	err = tag.DownloadAsset("frontend.zip", sourceFrontendZipPath)
	if err != nil {
		return core.InternalError("Error downloading frontend for version: " + tag.Tag + " " + err.Error())
	}

	readClose, err := zip.OpenReader(sourceFrontendZipPath)
	if err != nil {
		return core.InternalError("Error opening frontend.zip: " + err.Error())
	}

	for _, f := range readClose.File {
		name, _ := strings.CutPrefix(f.Name, "frontend/dist/")
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path.Join(targetFrontendPath, name), os.ModePerm)
			if err != nil {
				return core.InternalError("Error creating directory:" + err.Error())
			}
			continue
		}
		reader, _ := f.Open()

		file, err := os.Create(path.Join(targetFrontendPath, name))
		if err != nil {
			return core.InternalError("Error creating file " + file.Name() + err.Error())
		}
		defer file.Close()

		_, err = io.Copy(file, reader)
		if err != nil {
			return core.InternalError("Error copying file " + file.Name() + ": " + err.Error())
		} else {
			log.Success("Copied file " + file.Name())
		}
	}

	return core.StatusSuccess()
}
