package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestPermissionGroupUpdate(t *testing.T) {
	iamService := GetTestIAMService("iam")

	permissionGroupID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupUpdate",
		Abbr:        "PGU",
		Description: "test",
	})

	status := iamService.UpdatePermissionGroup(permissionGroupID, &iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupUpdate1",
		Abbr:        "PGUx",
		Description: "tset",
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	pgProfile, status := iamService.GetPermissionGroup(permissionGroupID)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if pgProfile == nil {
		t.Fatalf("No permission group found")
	}

	if pgProfile.Name != "PermissionGroupUpdate1" {
		t.Fatalf("Permission group name not updated")
	}

	if pgProfile.Abbr != "PGUx" {
		t.Fatalf("Permission group abbreviation not updated")
	}

	if pgProfile.Description != "tset" {
		t.Fatalf("Permission group description not updated")
	}
}
