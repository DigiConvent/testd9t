package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestAddPermissionToPermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	pg, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupAddPermission",
		Abbr:        "PGA",
		Description: "test",
		Parent:      getRootPermissionGroup(),
	})

	status := testService.AddPermissionToPermissionGroup(pg, "test")

	if !status.Err() && status.Message != "iam.permission_group.missing_permission" {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	status = testService.AddPermissionToPermissionGroup(nil, "test")

	if !status.Err() && status.Message != "iam.permission_group.missing_permission_group" {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	pgProfile, _ := testService.GetPermissionGroupProfile(pg)
	for _, permission := range pgProfile.PermissionGroup.Permissions {
		if permission.Name == "iam.user.write" {
			t.Fatal("Expected no permission, instead I got ", permission.Name)
		}
	}

	status = testService.AddPermissionToPermissionGroup(pg, "iam.user.write")
	if status.Err() {
		t.Fatal(status.Message)
	}

	pgProfile, _ = testService.GetPermissionGroupProfile(pg)
	found := false
	for _, permission := range pgProfile.Permissions {
		if permission.Name == "iam.user.write" {
			found = true
			break
		}
	}

	if !found {
		t.Log(pgProfile.Permissions)
		t.Fatal("Expected permission, instead I got ", pgProfile.Permissions)
	}
}
