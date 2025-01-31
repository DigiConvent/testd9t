package install

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core/file_repo"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
)

func GetFlavours(forVersion *sys_domain.Version) []string {
	if forVersion.Major == -1 {
		panic("Only call this for a specific version")
	}
	flavoursUrl := fmt.Sprint("https://raw.githubusercontent.com/DigiConvent/testd9t/refs/heads/", forVersion.String(), "/.meta/flavours.json")
	raw, err := file_repo.ReadJsonArray(flavoursUrl)

	if err != nil {
		fmt.Println(flavoursUrl)
		fmt.Println("Error fetching flavours:", err)
		return []string{}
	}

	flavours := make([]string, len(raw))
	for i, flavour := range raw {
		flavours[i] = flavour.(string)
	}

	return flavours
}
