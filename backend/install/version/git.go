package version

import (
	"context"

	"github.com/google/go-github/github"
)

func ListReleases() ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	releases, response, err := client.Repositories.ListReleases(context.Background(), "DigiConvent", "testd9t", nil)

	if err != nil || response.StatusCode != 200 {
		return nil, err
	}

	for i := 0; i < len(releases); i++ {
		iVersion := VersionFromString(*releases[i].TagName)
		for j := i + 1; j < len(releases); j++ {
			jVersion := VersionFromString(*releases[j].TagName)
			if iVersion == nil || jVersion == nil {
				continue
			}
			if jVersion.SmallerThan(iVersion) {
				releases[i], releases[j] = releases[j], releases[i]
			}
		}
	}

	return releases, nil
}
