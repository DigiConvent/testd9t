package update

import (
	"context"

	"github.com/google/go-github/github"
)

func Update() ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	releases, response, err := client.Repositories.ListReleases(context.Background(), "DigiConvent", "d9ttest", nil)

	if err != nil || response.StatusCode != 200 {
		return nil, err
	}

	return releases, nil
}
