package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestListUserPermissions(t *testing.T) {

	iamService := GetTestIAMService("iam")

	testUser := iam_domain.UserWrite{
		Emailaddress: "UserListPermissions@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := iamService.CreateUser(&testUser)

	permissions, _ := iamService.ListPermissions()
	// this is the permission group that the user is going to inherit fromt since its status will be a descendant of this permission group
	permissionGroup := iam_domain.PermissionGroupWrite{
		Name:        "TestUserListPermissions",
		Permissions: []string{permissions[0].Name},
		Abbr:        "TPG",
		Description: "Test Permission Group",
		Parent:      "",
	}
	permissionGroupID, _ := iamService.CreatePermissionGroup(&permissionGroup)

	userStatus := iam_domain.UserStatusWrite{
		Name:        "TestUserListPermissions",
		Abbr:        "TSLP",
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

	userPermissions, _ := iamService.ListUserPermissions(id)

	if len(userPermissions) != 1 {
		t.Errorf("User should have 1 permission")
	}

	if userPermissions[0].Name != permissions[0].Name {
		t.Errorf("User should have permission")
	}
}
