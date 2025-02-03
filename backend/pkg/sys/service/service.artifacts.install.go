package sys_service

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) InstallArtifacts(tag *sys_domain.ReleaseTag) *core.Status {
	log.Info("Installing artifacts for version " + tag.Tag)
	var err error
	homeFolder := "/home/testd9t/"
	tmpFolder := "/tmp/testd9t/"
	err = tag.DownloadAsset("main", homeFolder+"backend/main")
	if err != nil {
		return core.InternalError("Error downloading binary for version: " + tag.Tag + " " + err.Error())
	}

	err = tag.DownloadAsset("frontend.zip", tmpFolder+"frontend.zip")
	if err != nil {
		return core.InternalError("Error downloading frontend for version: " + tag.Tag + " " + err.Error())
	}

	err = exec.Command("chmod", "+x", homeFolder+"backend/main").Run()
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
			err := os.MkdirAll(homeFolder+"frontend/"+name, os.ModePerm)
			if err != nil {
				return core.InternalError("Error creating directory:" + err.Error())
			}
			continue
		}
		reader, _ := f.OpenRaw()

		file, err := os.Create(homeFolder + "frontend/" + name)
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

	os.Exit(0)

	return core.StatusSuccess()
}
