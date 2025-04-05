package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestDeleteUserStatus(t *testing.T) {
	iamService := GetTestIAMService("iam")

	rootId := uuid.MustParse(getRootPermissionGroup())
	id, status := iamService.CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "UserStatusDelete",
		Abbr:        "USD",
		Description: "test",
		Archived:    false,
		Parent:      &rootId,
	})

	if status.Err() {
		t.Errorf("CreateUserStatus() failed: %s", status.Message)
	}

	if id == nil {
		t.Fatalf("Expected a result")
	}

	status = iamService.DeleteUserStatus(id)

	if status.Err() {
		t.Errorf("DeleteUserStatus() failed: %s", status.Message)
	}

	userStatus, status := iamService.GetUserStatus(id)

	if !status.Err() {
		t.Errorf("GetUserStatus() failed: %s", status.Message)
	}

	if userStatus != nil {
		t.Errorf("Expected nil, got %v", userStatus)
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
