package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetUserStatus(t *testing.T) {
	iamService := GetTestIAMService("iam")

	root := uuid.MustParse(getRootPermissionGroup())
	id, status := iamService.CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "Test User Status",
		Abbr:        "TUS",
		Description: "testxs",
		Archived:    true,
		Parent:      &root,
	})

	if status.Err() {
		t.Errorf("CreateUserStatus() failed: %s", status.Message)
	}

	userStatusProfile, status := iamService.GetUserStatus(id)

	if status.Err() {
		t.Errorf("GetUserStatus() failed: %s", status.Message)
	}
	userStatus := userStatusProfile.UserStatus

	if userStatus == nil {
		t.Errorf("Expected a result, got %v", userStatus)
	}

	if userStatus.ID.String() != id.String() {
		t.Errorf("Expected %v, got %v", id, userStatus.ID.String())
	}

	if userStatus.Name != "Test User Status" {
		t.Errorf("Expected %s, got %s", "Test User Status", userStatus.Name)
	}

	if userStatus.Abbr != "TUS" {
		t.Errorf("Expected %s, got %s", "TUS", userStatus.Abbr)
	}

	if userStatus.Description != "testxs" {
		t.Errorf("Expected %s, got %s", "testxs", userStatus.Description)
	}

	if userStatus.Archived != true {
		t.Errorf("Expected %t, got %t", true, userStatus.Archived)
	}
}
