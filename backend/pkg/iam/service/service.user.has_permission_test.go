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
	// this is the permission group that the user is going to inherit from since its status will be a descendant of this permission group
	permissionGroup := iam_domain.PermissionGroupWrite{
		Name:        "TestPermissionGroupUserHasPermission",
		Permissions: []string{permissions[0].Name},
		Abbr:        "TPG",
		Description: "Test Permission Group",
		Parent:      "",
	}
	permissionGroupID, _ := iamService.CreatePermissionGroup(&permissionGroup)

	// make sure that the permission_group has this permission
	permissionGroupProfile, _ := iamService.GetPermissionGroupProfile(permissionGroupID)
	if permissionGroupProfile.Permissions[0].Name != permissions[0].Name {
		t.Errorf("Permission group should have permission " + permissions[0].Name)
	}

	userStatus := iam_domain.UserStatusWrite{
		Name:        "TestStatusUserHasPermission",
		Abbr:        "TS",
		Description: "Test Status",
		Archived:    false,
		Parent:      permissionGroupID,
	}
	statusID, _ := iamService.CreateUserStatus(&userStatus)

	if statusID == nil {
		t.Fatal("Expected a result")
	}

	status := iamService.AddUserStatus(&iam_domain.AddUserStatusToUser{
		UserID:   *id,
		StatusID: *statusID,
		When:     time.Now().Add(-2 * time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	permissionGroupProfile, _ = iamService.GetPermissionGroupProfile(statusID)
	parentPermissionGroupProfile, _ := iamService.GetPermissionGroupProfile(&permissionGroupProfile.Ancestors[0].ID)

	if parentPermissionGroupProfile.Permissions[0].Name != permissions[0].Name {
		t.Errorf("Permission group should have permission " + permissions[0].Name)
	}

	t.Log(permissionGroupProfile.Members)

	userProfile, _ := iamService.GetUserProfile(id)
	t.Log("User " + userProfile.User.Emailaddress + " has the following groups:")
	for _, group := range userProfile.Groups {
		t.Log(group.Name)
	}

	hasPermission := iamService.UserHasPermission(id, permissions[0].Name)
	if !hasPermission {
		userPermissions, _ := iamService.ListUserPermissions(id)
		t.Log("User has the following permissions:")
		for _, userPermission := range userPermissions {
			t.Log(userPermission.Name)
		}
		t.Errorf("User should have permission " + permissions[0].Name)
	}

	hasPermission = iamService.UserHasPermission(id, permissions[1].Name)
	if hasPermission {
		t.Errorf("User should have permission")
	}
}
