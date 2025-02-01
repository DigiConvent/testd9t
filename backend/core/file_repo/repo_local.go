package file_repo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type RepoLocal struct{}

func (m *RepoLocal) DownloadAsset(url string, path string) error {
	panic("unimplemented")
}

func (m *RepoLocal) Type() string {
	return "local"
}

func (m *RepoLocal) ReadRawFile(path string) ([]byte, error) {
	bytes, err := os.ReadFile(DevPath() + "/../" + path)
	if err != nil {
		fmt.Println("Could not read raw file: ", DevPath()+path)
		fmt.Println(err)
	}

	return bytes, nil
}

func NewRepoLocal() FileRepository {
	return &RepoLocal{}
}

func Dev() bool {
	return strings.HasPrefix(os.Args[0], "/tmp/go-build")
}

func DevPath() string {
	if Dev() {
		result, _ := exec.Command("pwd").Output()
		return strings.Replace(string(result), "\n", "", 1) + "/"
	}
	return ""
}
