package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetPermissionGroupProfile(t *testing.T) {
	testService := GetTestIAMService("iam")

	profileId, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupProfile",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
	})

	testUserId, _ := testService.CreateUser(&iam_domain.UserWrite{
		Email:       "PermissionGroupProfileTest@test.test",
		FirstName:   "Test",
		LastName:    "McTest",
		DateOfBirth: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	status := testService.AddUserToPermissionGroup(profileId, testUserId)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

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
}
