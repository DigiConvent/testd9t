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
		Parent:      getRootPermissionGroupUuid(),
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
		PermissionGroupWrite: iam_domain.PermissionGroupWrite{
			Name:        "PermissionGroupAddUserTest",
			Abbr:        "PGAUT",
			Description: "testxs",
			Parent:      &parsedId,
		},
		Archived: true,
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	if userStatus == nil {
		t.Fatal("Expected a result")
	}

	status = testService.AddUserToPermissionGroup(userStatus, user)

	if status.Err() {
		t.Fatal(status.Message)
	}

	// get profile and count users
	rootId := uuid.MustParse(getRootPermissionGroup())
	rootProfile, _ := testService.GetPermissionGroupProfile(&rootId)

	if rootProfile == nil {
		t.Fatal("Expected a result")
	}

	if len(rootProfile.Users) != 1 {
		t.Fatal("Expected 1 user, instead I got ", len(rootProfile.Users))
	}
}
