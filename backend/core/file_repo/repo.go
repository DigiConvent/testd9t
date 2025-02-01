package file_repo

type ReleaseTag struct {
	Tag            string `json:"tag"`
	MigrationFiles []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"migrations"`
	Assets []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"assets"`
}

type FileRepository interface {
	ReadRawFile(path string) ([]byte, error)
	DownloadAsset(url, path string) error
	Type() string
}
