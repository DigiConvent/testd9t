package sys_service

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func (s *SysService) InstallReleaseTag(tag *sys_domain.ReleaseTag) *core.Status {
	var err error
	homeFolder := "/home/testd9t"
	err = tag.DownloadAsset("main", homeFolder+"/backend/main")
	if err != nil {
		fmt.Println("Error downloading new version:", tag.Tag)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = tag.DownloadAsset("frontend.zip", "/tmp/frontend.zip")
	if err != nil {
		fmt.Println("Error downloading new version:", tag.Tag)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = exec.Command("chmod", "+x", homeFolder+"/backend/main").Run()
	if err != nil {
		return core.InternalError("Error setting permissions:" + err.Error())
	}

	readClose, err := zip.OpenReader("/tmp/frontend.zip")
	if err != nil {
		return core.InternalError("Error opening frontend.zip:" + err.Error())
	}

	for _, f := range readClose.File {
		name, _ := strings.CutPrefix(f.Name, "frontend/dist/")
		if f.FileInfo().IsDir() {
			err := os.MkdirAll("/tmp/testd9t/frontend/"+name, os.ModePerm)
			if err != nil {
				return core.InternalError("Error creating directory:" + err.Error())
			}
			continue
		}
		reader, _ := f.OpenRaw()

		file, err := os.Create("/tmp/testd9t/frontend/" + name)
		if err != nil {
			return core.InternalError("Error creating file:" + err.Error())
		}
		defer file.Close()

		io.Copy(file, reader)
	}

	os.Exit(0)

	return core.StatusSuccess()
}
