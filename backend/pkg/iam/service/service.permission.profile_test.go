package iam_service_test

import (
	"testing"
)

func TestGetPermissionProfile(t *testing.T) {
	iamService := GetTestIAMService("iam")

	permission, status := iamService.GetPermissionProfile("idam")
	if !status.Err() || permission != nil {
		t.Fatalf("GetPermissionProfile() succeeded where it should have failed")
	}

	permission, status = iamService.GetPermissionProfile("iam")

	if status.Err() {
		t.Fatalf("GetPermissionProfile() failed: %s", status.Message)
	}

	if permission == nil || permission.Permission == nil {
		t.Fatalf("GetPermissionProfile() failed: no permission found")
	}

	if permission.Permission.Name != "iam" {
		t.Fatalf("GetPermissionProfile() failed: wrong permission returned")
	}

	if permission.Descendants == nil {
		t.Fatalf("GetPermissionProfile() failed: no descendants found")
	}

	if len(permission.Descendants) == 0 {
		t.Fatalf("GetPermissionProfile() failed: no descendants found")
	}
}
