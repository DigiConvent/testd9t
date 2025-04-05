package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	id, status := iamService.CreateUserRole(&iam_domain.UserRoleWrite{
		Name:        "Test User Status",
		Abbr:        "TUS",
		Description: "testxs",
	})

	if status.Err() {
		t.Errorf("CreateUserRole() failed: %s", status.Message)
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

	if UserRole.Name != "Test User Status" {
		t.Errorf("Expected %s, got %s", "Test User Status", UserRole.Name)
	}

	if UserRole.Abbr != "TUS" {
		t.Errorf("Expected %s, got %s", "TUS", UserRole.Abbr)
	}

	if UserRole.Description != "testxs" {
		t.Errorf("Expected %s, got %s", "testxs", UserRole.Description)
	}

}
