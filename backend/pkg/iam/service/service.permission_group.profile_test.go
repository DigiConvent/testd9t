package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetPermissionGroupProfile(t *testing.T) {
	testService := GetTestIAMService("iam")

	permissions, _ := testService.ListPermissions()

	profileId, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupProfile",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
		Permissions: []string{permissions[0].Name, permissions[1].Name, permissions[2].Name},
		Parent:      getRootPermissionGroup(),
	})

	permissionGroupProfile, status := testService.GetPermissionGroupProfile(profileId)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if permissionGroupProfile == nil {
		t.Fatalf("No permission group found")
	}

	if permissionGroupProfile.PermissionGroup.Name != "PermissionGroupProfile" {
		t.Fatalf("Expected PermissionGroupProfile, instead got %v", permissionGroupProfile.PermissionGroup.Name)
	}

	if permissionGroupProfile.PermissionGroup.Abbr != "PG" {
		t.Fatalf("Expected PG, instead got %v", permissionGroupProfile.PermissionGroup.Abbr)
	}

	if permissionGroupProfile.PermissionGroup.Description != "test" {
		t.Fatalf("Expected test, instead got %v", permissionGroupProfile.PermissionGroup.Description)
	}

	if !permissionGroupProfile.PermissionGroup.IsGroup {
		t.Fatalf("Expected true, instead got %v", permissionGroupProfile.PermissionGroup.IsGroup)
	}

	if permissionGroupProfile.PermissionGroup.IsNode {
		t.Fatalf("Expected false, instead got %v", permissionGroupProfile.PermissionGroup.IsNode)
	}

	if permissionGroupProfile.PermissionGroup.ID != *profileId {
		t.Fatalf("Expected %v, instead got %v", profileId, permissionGroupProfile.PermissionGroup.ID)
	}

	if len(permissionGroupProfile.Permissions) != 3 {
		t.Fatalf("Expected 3, instead got %v", len(permissionGroupProfile.Permissions))
	}
}
