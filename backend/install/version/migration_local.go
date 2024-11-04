package version

import (
	"fmt"
	"os"
)

func (l *localMigrationsMigrator) ListMigrationFiles(version *SemVer) map[string]string {
	result := make(map[string]string)
	dirPath := fmt.Sprintf("%s/%s", os.Getenv("MIGRATIONS_DIR"), version.String())
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", dirPath, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return nil
		}

		result[file.Name()] = string(content)
	}

	return result
}

func (l *localMigrationsMigrator) ListMigrationVersions() []SemVer {
	result := make([]SemVer, 0)
	files, err := os.ReadDir(os.Getenv("MIGRATIONS_DIR"))
	if err != nil {
		fmt.Println("Error reading directory:", os.Getenv("MIGRATIONS_DIR"), err)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		ver := VersionFromString(file.Name())
		if ver == nil {
			continue
		}

		result = append(result, *VersionFromString(ver.String()))
	}

	Sort(result, true)

	return result
}
