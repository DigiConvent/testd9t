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
	}

	id, _ := iamService.CreateUserRole(testUserRole)

	if id == nil {
		t.Fatal("Expected a result")
	}

	testUserRole.Name = "UserRoleUpdateUpdated"
	testUserRole.Abbr = "USUU"
	testUserRole.Description = "testx"

	status := iamService.UpdateUserRole(id, testUserRole)

	if status.Err() {
		t.Fatal(status.Message)
	}

	UserRoleProfile, status := iamService.GetUserRole(id)

	if status.Err() {
		t.Errorf("GetUserRole() failed: %s", status.Message)
	}

	UserRole := UserRoleProfile.UserRole

	if UserRole == nil {
		t.Errorf("Expected a result, got %v", UserRole)
	}

	if UserRole.ID.String() != id.String() {
		t.Errorf("Expected %v, got %v", id, UserRole.ID.String())
	}

	if UserRole.Name != "UserRoleUpdateUpdated" {
		t.Errorf("Expected %s, got %s", "UserRoleUpdateUpdated", UserRole.Name)
	}

	if UserRole.Abbr != "USUU" {
		t.Errorf("Expected %s, got %s", "USUU", UserRole.Abbr)
	}

	if UserRole.Description != "testx" {
		t.Errorf("Expected %s, got %s", "testx", UserRole.Description)
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
