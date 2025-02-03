package sys_domain

import (
	"errors"

	"github.com/DigiConvent/testd9t/core/file_repo"
)

type ReleaseTag struct {
	Tag        string   `json:"tag"`
	Assets     []string `json:"assets"`
	Migrations []string `json:"migrations"`
}

func (tag *ReleaseTag) AssetURL(name string) string {
	for _, fileName := range tag.Assets {
		if fileName == name {
			return "https://github.com/" + file_repo.GHUser + "/" + file_repo.GHRepo + "/releases/download/" + tag.Tag + "/" + name
		}
	}
	return ""
}

func (tag *ReleaseTag) MigrationURL(name string) string {
	for _, fileName := range tag.Migrations {
		if fileName == name {
			return "https://github.com/" + file_repo.GHUser + "/" + file_repo.GHRepo + "/refs/tags/" + tag.Tag + "/data/migrations/" + tag.Tag + "/" + name
		}
	}
	return ""
}

func (tag *ReleaseTag) DownloadAsset(name, path string) error {
	url := tag.AssetURL(name)
	if url == "" {
		return errors.New("Asset not found: " + name)
	}
	return file_repo.NewRepoRemote().DownloadAsset(url, path)
}
