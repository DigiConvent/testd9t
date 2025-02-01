package file_repo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix string = "https://raw.githubusercontent.com/DigiConvent/testd9t/refs/heads/main/"

const (
	GHUser = "DigiConvent"
	GHRepo = "testd9t"
)

func (mr *RepoRemote) ReadRawFile(filePath string) ([]byte, error) {
	if filePath[0] == '/' {
		filePath = filePath[1:]
	}

	if !strings.HasPrefix(filePath, "http") {
		filePath = prefix + filePath
	}

	resp, err := http.Get(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", filePath, resp.Status)
	}

	return io.ReadAll(resp.Body)
}

type RepoRemote struct{}

func (mr *RepoRemote) DownloadAsset(url, filepath string) error {
	os.Remove(filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	return err
}

func (mr *RepoRemote) Type() string {
	return "remote"
}

func NewRepoRemote() FileRepository {
	return &RepoRemote{}
}
