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

	if userRole.Name != "Test User Status" {
		t.Errorf("Expected %s, got %s", "Test User Status", userRole.Name)
	}

	if userRole.Abbr != "TUS" {
		t.Errorf("Expected %s, got %s", "TUS", userRole.Abbr)
	}

	if userRole.Description != "testxs" {
		t.Errorf("Expected %s, got %s", "testxs", userRole.Description)
	}
}
