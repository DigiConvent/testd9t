package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (migrator *remoteMigrationsMigrator) ListMigrationVersions() []SemVer {
	result := make([]SemVer, 0)
	if os.Getenv("GH_TOKEN") != "" {

	}
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/backend/data/db/migrations/", RepositoryUser, RepositoryName)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GH_TOKEN")))
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("Could not create request:", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error fetching migration versions:", response.Status)
		fmt.Println("URL:", url)
		return nil
	}

	var responseJson []map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseJson); err != nil {
		contents, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		} else {
			fmt.Println(string(contents))
		}

		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	for i := 0; i < len(responseJson); i++ {
		if responseJson[i]["type"] == "dir" {
			result = append(result, *VersionFromString(responseJson[i]["name"].(string)))
		}
	}

	Sort(result, true)

	return result
}

func (migrator *remoteMigrationsMigrator) ListMigrationFiles(version *SemVer) map[string]string {
	result := make(map[string]string)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/backend/data/db/migrations/%s/", RepositoryUser, RepositoryName, version.String())

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not create request:", err)
		return nil
	}
	if response.StatusCode != 200 {
		fmt.Println("Error fetching migration files:", response.Status)
		return nil
	}

	defer response.Body.Close()

	var responseJson []map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseJson); err != nil {
		fmt.Println("Whats going on")
		fmt.Println("Error decoding JSON:", err)
		contents, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil
		}
		fmt.Println(string(contents))
		return nil
	}

	for i := 0; i < len(responseJson); i++ {
		if responseJson[i]["type"] == "file" {
			downloadResponse, err := http.Get(responseJson[i]["download_url"].(string))
			if err != nil {
				fmt.Println("Could not create request:", err)
				return nil
			}
			defer downloadResponse.Body.Close()

			contents, err := io.ReadAll(downloadResponse.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return nil
			}
			result[responseJson[i]["name"].(string)] = string(contents)
		}
	}

	return result
}
