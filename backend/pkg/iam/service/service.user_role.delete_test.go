package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestDeleteUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	id, status := iamService.CreateUserRole(&iam_domain.UserRoleWrite{
		Name:        "UserRoleDelete",
		Abbr:        "USD",
		Description: "test",
	})

	if status.Err() {
		t.Errorf("CreateUserRole() failed: %s", status.Message)
	}

	if id == nil {
		t.Fatalf("Expected a result")
	}

	status = iamService.DeleteUserRole(id)

	if status.Err() {
		t.Errorf("DeleteUserRole() failed: %s", status.Message)
	}

	UserRole, status := iamService.GetUserRole(id)

	if !status.Err() {
		t.Errorf("GetUserRole() failed: %s", status.Message)
	}

	if UserRole != nil {
		t.Errorf("Expected nil, got %v", UserRole)
	}

	pg, status := iamService.GetPermissionGroup(id)

	if status == nil {
		t.Fatalf("Status cannot be nil")
	}

	if !status.Err() {
		t.Errorf("GetPermissionGroup() failed: %s", status.Message)
	}

	if pg != nil {
		t.Errorf("Expected nil, got %v", pg)
	}
}
