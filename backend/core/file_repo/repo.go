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
	GetRawFile(path string) ([]byte, error)
	Type() string
}
