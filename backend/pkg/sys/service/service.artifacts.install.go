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

func (s *SysService) InstallArtifacts(tag *sys_domain.ReleaseTag) *core.Status {
	log.Info("Installing artifacts for version " + tag.Tag)
	var err error
	tmpFolder := "/tmp/testd9t/"
	err = tag.DownloadAsset("main", constants.HOME_PATH+"backend/main")
	if err != nil {
		return core.InternalError("Error downloading binary for version: " + tag.Tag + " " + err.Error())
	}

	err = tag.DownloadAsset("frontend.zip", tmpFolder+"frontend.zip")
	if err != nil {
		return core.InternalError("Error downloading frontend for version: " + tag.Tag + " " + err.Error())
	}

	err = exec.Command("chmod", "+x", constants.HOME_PATH+"backend/main").Run()
	if err != nil {
		return core.InternalError("Error setting permissions for binary: " + err.Error())
	}

	readClose, err := zip.OpenReader(tmpFolder + "frontend.zip")
	if err != nil {
		return core.InternalError("Error opening frontend.zip: " + err.Error())
	}

	for _, f := range readClose.File {
		name, _ := strings.CutPrefix(f.Name, "frontend/dist/")
		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path.Join(constants.HOME_PATH, "frontend", name), os.ModePerm)
			if err != nil {
				return core.InternalError("Error creating directory:" + err.Error())
			}
			continue
		}
		reader, _ := f.Open()

		file, err := os.Create(path.Join(constants.HOME_PATH, "frontend", name))
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

	// give correct rights
	exec.Command("chown", "-R", "testd9t:testd9t", constants.HOME_PATH+"frontend").Run()

	os.Exit(0)

	return core.StatusSuccess()
}
