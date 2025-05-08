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
	// this is the permission group that the user is going to inherit from since its status will be a descendant of this permission group
	permissionGroup := iam_domain.PermissionGroupWrite{
		Name:        "TestUserListPermissions",
		Permissions: []string{permissions[0].Name},
		Abbr:        "TPG",
		Description: "Test Permission Group",
		Parent:      getRootPermissionGroupUuid(),
	}
	permissionGroupID, _ := iamService.CreatePermissionGroup(&permissionGroup)

	userStatus := iam_domain.UserStatusWrite{
		PermissionGroupWrite: iam_domain.PermissionGroupWrite{
			Name:        "TestUserListPermissions",
			Abbr:        "TSLP",
			Description: "Test Status",
			Parent:      permissionGroupID,
		},
		Archived: false,
	}
	statusID, _ := iamService.CreateUserStatus(&userStatus)

	if statusID == nil {
		t.Fatal("Expected a result")
	}

	iamService.AddUserToUserStatus(&iam_domain.UserBecameStatusWrite{
		User:       *id,
		UserStatus: *statusID,
		Start:      time.Now().Add(-2 * time.Hour),
	})

	userPermissions, _ := iamService.ListUserPermissions(id)

	if len(userPermissions) != 1 {
		t.Errorf("User should have 1 permission")
	}

	// if userPermissions[0].Name != permissions[0].Name {
	// 	t.Errorf("User should have permission")
	// }
}
