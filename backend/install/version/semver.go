package version

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
