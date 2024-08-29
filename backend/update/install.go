package update

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
)

func Install(release *github.RepositoryRelease) error {
	// download release to /tmp
	client := github.NewClient(nil)

	for i, e := range release.Assets {
		fmt.Println(*release.Assets[i].Name)

		out, err := os.Create("/tmp/" + *e.Name)
		if err != nil {
			return err
		}
		defer out.Close()

		// get file
		read, _, err := client.Repositories.DownloadReleaseAsset(context.Background(), "DigiConvent", "d9ttest", *e.ID)

		if err != nil {
			return err
		}
		defer read.Close()

		// write file
		_, err = out.ReadFrom(read)
		if err != nil {
			return err
		}
	}

	return nil
}
