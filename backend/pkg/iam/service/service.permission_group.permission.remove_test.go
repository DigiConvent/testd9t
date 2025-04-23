package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestRemovePermissionFromPermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	status := testService.RemovePermissionFromPermissionGroup(nil, "test")

	if !status.Err() && status.Message != "iam.permission_group.missing_permission_group" {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	pg, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupRemovePermission",
		Abbr:        "PG",
		Description: "test",
		Parent:      getRootPermissionGroup(),
	})

	status = testService.RemovePermissionFromPermissionGroup(pg, "test")
	if !status.Err() && status.Message != "iam.permission_group.missing_permission" {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	pgProfile, _ := testService.GetPermissionGroupProfile(pg)

	if len(pgProfile.Permissions) != 0 {
		t.Fatal("Expected 0 permission, instead I got ", len(pgProfile.Permissions))
	}

	testService.AddPermissionToPermissionGroup(pg, "iam.user.write")

	pgProfile, _ = testService.GetPermissionGroupProfile(pg)

	if len(pgProfile.Permissions) != 1 {
		t.Fatal("Expected 1 permission, instead I got ", len(pgProfile.Permissions))
	}

	status = testService.RemovePermissionFromPermissionGroup(pg, "iam.user.write")

	if status.Err() {
		t.Fatal(status.Message)
	}

	if pgProfile.Permissions[0].Name != "iam.user.write" {
		t.Fatal("Expected iam.user.write, instead I got ", pgProfile.Permissions[0].Name)
	}

	testService.RemovePermissionFromPermissionGroup(pg, "iam.user.write")

	pgProfile, _ = testService.GetPermissionGroupProfile(pg)

	if len(pgProfile.Permissions) != 0 {
		t.Fatal("Expected 0 permissions, instead I got ", len(pgProfile.Permissions))
	}
}
