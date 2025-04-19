package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestListPermissionGroups(t *testing.T) {
	testService := GetTestIAMService("iam")

	testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupList1",
		Abbr:        "PG",
		Description: "test",
		Parent:      getRootPermissionGroup(),
	})
	testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupList2",
		Abbr:        "PG",
		Description: "test",
		Parent:      getRootPermissionGroup(),
	})

	permissionGroups, status := testService.ListPermissionGroups()

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if len(permissionGroups) == 0 {
		t.Fatalf("No permission groups found")
	}

	if len(permissionGroups) < 2 {
		t.Fatalf("Expected at least 2 permission groups, instead got %v", len(permissionGroups))
	}

	found1, found2 := false, false
	for _, permissionGroup := range permissionGroups {
		if permissionGroup.Name == "PermissionGroupList1" {
			found1 = true
		}
		if permissionGroup.Name == "PermissionGroupList2" {
			found2 = true
		}
	}

	if !found1 {
		t.Fatalf("Expected PermissionGroupList1, not found")
	}

	if !found2 {
		t.Fatalf("Expected PermissionGroupList2, not found")
	}
}
