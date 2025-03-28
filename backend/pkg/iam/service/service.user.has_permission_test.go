package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestUserHasPermission(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser := iam_domain.UserWrite{
		Emailaddress: "UserHasPermission@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := iamService.CreateUser(&testUser)

	permissions, _ := iamService.ListPermissions()
	// this is the permission group that the user is going to inherit fromt since its status will be a descendant of this permission group
	permissionGroup := iam_domain.PermissionGroupWrite{
		Name:        "TestPermissionGroupUserHasPermission",
		Permissions: []string{permissions[0].Name},
		Abbr:        "TPG",
		Description: "Test Permission Group",
		Parent:      "",
	}
	permissionGroupID, _ := iamService.CreatePermissionGroup(&permissionGroup)

	userStatus := iam_domain.UserStatusWrite{
		Name:        "TestStatusUserHasPermission",
		Abbr:        "TS",
		Description: "Test Status",
		Archived:    false,
	}
	statusID, _ := iamService.CreateUserStatus(&userStatus)

	iamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		ID:     statusID,
		Parent: permissionGroupID,
	})

	iamService.AddUserStatus(&iam_domain.AddUserStatusToUser{
		UserID:      *id,
		StatusID:    *statusID,
		Description: "Test Status",
		When:        time.Now().Add(-2 * time.Hour),
	})

	hasPermission := iamService.UserHasPermission(id, permissions[0].Name)
	if !hasPermission {
		t.Errorf("User should have permission " + permissions[0].Name)
	}

	hasPermission = iamService.UserHasPermission(id, permissions[1].Name)
	if hasPermission {
		t.Errorf("User should have permission")
	}
}
