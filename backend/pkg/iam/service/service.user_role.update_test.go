package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestUpdateUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUserRole := &iam_domain.UserRoleWrite{
		Name:        "UserRoleUpdate",
		Abbr:        "USU",
		Description: "test",
		Parent:      getRootPermissionGroupUuid(),
	}

	id, _ := iamService.CreateUserRole(testUserRole)

	if id == nil {
		t.Fatal("Expected a result")
	}

	testUserRole.Name = ""

	status := iamService.UpdateUserRole(id, testUserRole)

	if !status.Err() && status.Message != "iam.user_role.update.invalid_name" {
		t.Fatalf("expected error, got %v", status.Message)
	}

	testUserRole.Name = "UserRoleUpdateUpdated"
	testUserRole.Abbr = "USUU"
	testUserRole.Description = "testx"
	testUserRole.Parent = nil

	status = iamService.UpdateUserRole(id, testUserRole)

	if !status.Err() && status.Message != "iam.user_role.update.invalid_parent" {
		t.Fatalf("expected error, got %v", status.Message)
	}

	testUserRole.Parent = getRootPermissionGroupUuid()

	status = iamService.UpdateUserRole(id, testUserRole)

	if status.Err() {
		t.Errorf("UpdateUserRole() failed: %s", status.Message)
	}

	userRole, status := iamService.GetUserRole(id)

	if status.Err() {
		t.Errorf("GetUserRole() failed: %s", status.Message)
	}

	if userRole == nil {
		t.Errorf("Expected a result, got %v", userRole)
	}

	if userRole.Id.String() != id.String() {
		t.Errorf("Expected %v, got %v", id, userRole.Id.String())
	}

	if userRole.Name != testUserRole.Name {
		t.Errorf("Expected %s, got %s", testUserRole.Name, userRole.Name)
	}

	if userRole.Abbr != testUserRole.Abbr {
		t.Errorf("Expected %s, got %s", testUserRole.Abbr, userRole.Abbr)
	}

	if userRole.Description != "testx" {
		t.Errorf("Expected %s, got %s", "testx", userRole.Description)
	}

	pg, _ := iamService.GetPermissionGroup(id)

	if pg == nil {
		t.Fatal("Expected a result")
	}

	if pg.Name != testUserRole.Name {
		t.Errorf("Expected %s, got %s", testUserRole.Name, pg.Name)
	}

	if pg.Abbr != testUserRole.Abbr {
		t.Errorf("Expected %s, got %s", testUserRole.Abbr, pg.Abbr)
	}

	if pg.Description != testUserRole.Description {
		t.Errorf("Expected %s, got %s", testUserRole.Description, pg.Description)
	}
}
