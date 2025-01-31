package file_repo

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const prefix string = "https://raw.githubusercontent.com/DigiConvent/testd9t/refs/heads/main/"

const (
	GHUser = "DigiConvent"
	GHRepo = "testd9t"
)

func (mr *RepoRemote) GetRawFile(filePath string) ([]byte, error) {
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

func (mr *RepoRemote) Type() string {
	return "remote"
}

func NewRepoRemote() FileRepository {
	return &RepoRemote{}
}
