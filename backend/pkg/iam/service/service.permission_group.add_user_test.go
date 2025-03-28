package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
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
	})

	if user == nil {
		t.Fatal("Expected a result")
	}

	testService.AddUserToPermissionGroup(pg, user)

	id := getRootPermissionGroup()
	parsedId, _ := uuid.Parse(id)

	userStatus, status := testService.CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "PermissionGroupAddUserTest",
		Abbr:        "PGAUT",
		Description: "testxs",
		Archived:    true,
		Parent:      &parsedId,
	})

	if status.Err() {
		t.Log("Cannot create user status")
		t.Fatal(status.Message)
	}

	if userStatus == nil {
		t.Fatal("Expected a result")
	}

	status = testService.AddUserToPermissionGroup(userStatus, user)

	if !status.Err() {
		t.Log("Cannot add user to permission group")
		t.Fatal(status.Message)
	}
}
