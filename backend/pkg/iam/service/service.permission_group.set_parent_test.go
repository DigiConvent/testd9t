package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestSetPermissionGroupParent(t *testing.T) {
	iamService := GetTestIAMService("iam")

	permissionGroupGrandParentID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupGrandParent",
		Abbr:        "PGP",
		Description: "test",
		Parent:      getRootPermissionGroupUuid(),
	})
	permissionGroupChildID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupParentSetChild",
		Abbr:        "PGP",
		Description: "test",
		Parent:      getRootPermissionGroupUuid(),
	})

	permissionGroupParentID, _ := iamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupParent",
		Abbr:        "PGP",
		Description: "test",
		Parent:      getRootPermissionGroupUuid(),
	})

	status := iamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		Id:     permissionGroupChildID,
		Parent: permissionGroupParentID,
	})

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	status = iamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		Id:     permissionGroupParentID,
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

	if len(pgProfile.Ancestors) != 4 {
		t.Fatalf("Expected 4 permission group, instead got %v", len(pgProfile.Ancestors))
	}
}
