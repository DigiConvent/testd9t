package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestAddUserToPermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	pg, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupAddUser",
		Abbr:        "PGA",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
		Parent:      getRootPermissionGroup(),
	})

	if pg == nil {
		t.Fatal("Expected a result")
	}

	user, _ := testService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "PermissionGroupAddUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
		DateOfBirth:  time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	if user == nil {
		t.Fatal("Expected a result")
	}

	status := testService.AddUserToPermissionGroup(pg, user)

	if status.Err() {
		t.Fatal(status.Message)
	}

	// check if a user can be added to a generated permission group

	userStatus, _ := testService.CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "PermissionGroupAddUserTest",
		Abbr:        "PGAUT",
		Description: "testxs",
		Archived:    true,
	})

	if userStatus == nil {
		t.Fatal("Expected a result")
	}

	status = testService.AddUserToPermissionGroup(userStatus, user)

	if !status.Err() {
		t.Fatal(status.Message)
	}
}
