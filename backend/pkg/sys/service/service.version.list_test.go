package sys_service_test

import "testing"

func TestListVersions(t *testing.T) {
	sysService := GetTestSysService("sys")

	versions, status := sysService.ListReleaseTags()

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if len(versions) == 0 {
		t.Fatalf("No versions found")
	}
}
