package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestSetPermissionGroupParent(t *testing.T) {
	iamService := GetTestIAMService("iam")

	permissionGroupChildID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupParentSetChild",
		Abbr:        "PGP",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
	})

	permissionGroupParentID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupParent",
		Abbr:        "PGP",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
	})
	permissionGroupGrandParentID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupGrandParent",
		Abbr:        "PGP",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
	})

	status := iamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		ID:     permissionGroupChildID,
		Parent: permissionGroupParentID,
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	status = iamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		ID:     permissionGroupParentID,
		Parent: permissionGroupGrandParentID,
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	pgProfile, status := iamService.GetPermissionGroupProfile(permissionGroupChildID)

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if pgProfile == nil {
		t.Fatalf("No permission group found")
	}

	if len(pgProfile.PermissionGroups) != 3 {
		t.Fatalf("Expected 1 permission group, instead got %v", len(pgProfile.PermissionGroups))
	}
}
