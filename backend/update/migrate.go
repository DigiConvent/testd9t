package update

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type SemVer struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type D9TRelease struct {
	Version     SemVer `json:"version"`
	Description string `json:"description"`
}

func (v *SemVer) String() string {
	patch := fmt.Sprintf("%d", v.Patch)
	if v.Patch == -1 {
		patch = "x"
	}
	return fmt.Sprintf("%d.%d.%s", v.Major, v.Minor, patch)
}

func (v *SemVer) TimeOfDownload() *time.Time {
	file, err := os.Stat("releases/" + v.String())
	if err != nil {
		return nil
	}
	var dTime = file.ModTime()
	return &dTime
}

func VersionFromString(version string) *SemVer {
	if !strings.Contains(version, ".") {
		return nil
	}
	if strings.Count(version, ".") != 2 {
		return nil
	}

	parts := strings.Split(version, ".")
	var major, minor, patch int
	var err error
	if major, err = strconv.Atoi(parts[0]); err != nil {
		return nil
	}
	if minor, err = strconv.Atoi(parts[1]); err != nil {
		return nil
	}
	if patch, err = strconv.Atoi(parts[2]); err != nil {
		return nil
	}

	return &SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (v *SemVer) MigrationFolderExists() bool {
	path := os.Getenv("CONFIG_DIR") + "migrations/" + v.String()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Could not find migration folder for version", v.String(), ": ", path)
		return false
	}
	return info.IsDir()
}

func (v *SemVer) MigrationScripts() []string {
	path := os.Getenv("CONFIG_DIR") + "migrations/" + v.String() + "/"
	scriptPaths, err := os.ReadDir(path)

	var scripts = make([]string, 0)

	if err != nil && v.Patch != -1 {
		fmt.Println(err.Error())
		fmt.Println("Could not read migration folder for version", v.String(), ": ", path)
		return scripts
	}

	for _, script := range scriptPaths {
		if script.IsDir() {
			continue
		}

		contents, err := os.ReadFile(path + script.Name())
		if err != nil {
			fmt.Println("Could not read migration script", script.Name(), "for version", v.String(), ": ", path)
			continue
		}
		scripts = append(scripts, string(contents))
	}

	if v.Patch == -1 {
		return scripts
	}

	versionData := *v
	versionData.Patch = -1

	dataScripts := versionData.MigrationScripts()

	scripts = append(scripts, dataScripts...)

	return scripts
}

func (v *SemVer) Equals(other *SemVer) bool {
	return v.Major == other.Major && v.Minor == other.Minor && v.Patch == other.Patch
}

func (v *SemVer) SmallerThan(other *SemVer) bool {
	if v.Major < other.Major {
		return true
	}
	if v.Major == other.Major && v.Minor < other.Minor {
		return true
	}
	if v.Major == other.Major && v.Minor == other.Minor && v.Patch < other.Patch {
		return true
	}
	return false
}

func (end *SemVer) MigrationsBetween(start *SemVer, versions *[]SemVer) []SemVer {
	if versions == nil {
		return nil
	} else {
		var migrations []SemVer
		for _, version := range *versions {
			if start.SmallerThan(&version) && version.SmallerThan(end) {
				migrations = append(migrations, version)
			}
		}
		return migrations
	}
}

func ListMigrations(from, to *SemVer) []SemVer {
	path := os.Getenv("CONFIG_DIR") + "migrations/"
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Could not read migrations folder: ", err.Error())
		return nil
	}
	var migrations []SemVer
	for _, file := range files {
		if file.IsDir() {
			version := VersionFromString(file.Name())
			if version == nil {
				continue
			}
			if from != nil && version.SmallerThan(from) {
				continue
			}
			if to != nil && to.SmallerThan(version) {
				continue
			}
			migrations = append(migrations, *version)
		}
	}
	return migrations
}

func Sort(version []SemVer, asc bool) {
	for i := 0; i < len(version); i++ {
		for j := i + 1; j < len(version); j++ {
			if asc {
				if version[j].SmallerThan(&version[i]) {
					version[i], version[j] = version[j], version[i]
				}
			} else {
				if version[i].SmallerThan(&version[j]) {
					version[i], version[j] = version[j], version[i]
				}
			}
		}
	}
}
