package sys_domain

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

var ProgramVersion string = "dev"
var DatabasePath string = "/tmp/d9t/db/"
var CompiledAt string = ""

func DevPath() string {
	if ProgramVersion == "dev" {
		result, _ := exec.Command("pwd").Output()
		return strings.Replace(string(result), "\n", "", 1) + "/"
	}
	return ""
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}
type Migration struct {
	Version   Version `json:"version"`
	Migration string  `json:"migration"`
}

func (m *Migration) SmallerThan(other *Migration) bool {
	return m.Version.SmallerThan(&other.Version) && m.Migration < other.Migration
}

func (m *Migration) Equals(other *Migration) bool {
	return m.Version.Equals(&other.Version) && m.Migration == other.Migration
}

func (v *Version) String() string {
	patch := fmt.Sprintf("%d", v.Patch)
	return fmt.Sprintf("%d.%d.%s", v.Major, v.Minor, patch)
}

func VersionFromString(version string) *Version {
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

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (v *Version) SmallerThan(other *Version) bool {
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

func (v *Version) SmallerThanOrEqual(other *Version) bool {
	if v.SmallerThan(other) || v.Equals(other) {
		return true
	}
	return false
}

func (v *Version) Equals(other *Version) bool {
	return v.Major == other.Major && v.Minor == other.Minor && v.Patch == other.Patch
}

func Sort(version []Version, asc bool) {
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
