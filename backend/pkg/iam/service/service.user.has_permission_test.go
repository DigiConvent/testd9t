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

	permission := "iam.user"
	permissionNotImpliedByParentPermission := "iam.user_status.read"

	// this is the permission group that the user is going to inherit from since its status will be a descendant of this permission group
	permissionGroup := iam_domain.PermissionGroupWrite{
		Name:        "TestPermissionGroupUserHasPermission",
		Permissions: []string{permission},
		Abbr:        "TPG",
		Description: "Test Permission Group",
		Parent:      getRootPermissionGroup(),
	}
	permissionGroupID, _ := iamService.CreatePermissionGroup(&permissionGroup)

	// make sure that the permission_group has this permission
	permissionGroupProfile, _ := iamService.GetPermissionGroupProfile(permissionGroupID)
	if permissionGroupProfile.Permissions[0].Name != permission {
		t.Fatalf("Permission group should have permission " + permission)
	}

	userStatus := iam_domain.UserStatusWrite{
		Name:        "TestStatusUserHasPermission",
		Abbr:        "TS",
		Description: "Test Status",
		Archived:    false,
		Parent:      permissionGroupID,
	}

	statusID, status := iamService.CreateUserStatus(&userStatus)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if statusID == nil {
		t.Fatal("Expected a result")
	}

	status = iamService.AddUserBecameStatus(&iam_domain.UserBecameStatusWrite{
		User:       *id,
		UserStatus: *statusID,
		Start:      time.Now().Add(-2 * time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	permissionGroupProfile, _ = iamService.GetPermissionGroupProfile(statusID)
	parentPermissionGroupProfile, status := iamService.GetPermissionGroupProfile(permissionGroupProfile.PermissionGroup.Parent)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if len(parentPermissionGroupProfile.Permissions) <= 0 {
		t.Fatalf("Permission group should have permissions ")
	}

	if parentPermissionGroupProfile.Permissions[0].Name != permission {
		t.Fatalf("Permission group should have permission " + permission)
	}

	_, status = iamService.GetUserProfile(id)

	if status.Err() {
		t.Fatalf(status.Message)
	}

	hasPermission := iamService.UserHasPermission(id, permission)
	if !hasPermission {
		userPermissions, _ := iamService.ListUserPermissions(id)
		t.Log("User has the following permissions:")
		for _, userPermission := range userPermissions {
			t.Log(userPermission.Name)
		}
		t.Fatalf("But user should have permission " + permission)
	}

	hasPermission = iamService.UserHasPermission(id, permissionNotImpliedByParentPermission)
	if hasPermission {
		t.Fatalf("User should not have permission " + permissionNotImpliedByParentPermission + " since it is not implied by " + permission)
	}
}
