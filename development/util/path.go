package util

import (
	"os/exec"
	"strings"
)

func GetPath() string {
	result, _ := exec.Command("pwd").Output()
	devPath := strings.Replace(string(result), "\n", "", 1) + "/"

	if strings.Contains(devPath, "testd9t") {
		if strings.Count(devPath, "testd9t") != 1 {
			panic("There should be only one testd9t folder and that is the project root folder")
		}
		if strings.HasSuffix(devPath, "testd9t/") {
			return devPath
		}

		segments := strings.Split(devPath, "/")
		for i := 0; i < len(segments); i++ {
			if segments[i] == "testd9t" {
				return strings.Join(segments[:i+1], "/") + "/"
			}
		}
	}

	panic("Invalid path, execute from anywhere inside the testd9t folder")
}

func GetFrontendPath() string {
	return GetPath() + "frontend/"
}

func GetBackendPath() string {
	return GetPath() + "backend/"
}

func GetSharedPath() string {
	return GetPath() + "shared/"
}
