package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestDeletePermission(t *testing.T) {
	service := GetTestIAMService("iam")

	permissionName := "some.test.permission"
	service.CreatePermission(&iam_domain.PermissionWrite{
		Name:        permissionName,
		Description: "test",
		Meta:        "",
	})

	status := service.DeletePermission(permissionName)
	if status != nil && status.Err() {
		t.Errorf("Error deleting permission: " + status.Message)
	}

	if status.Code != 204 {
		t.Errorf("Expected 204, got %v", status.Code)
	}

	permissions, _ := service.ListPermissions()

	for _, permission := range permissions {
		if permission.Name == permissionName {
			t.Errorf("Expected permission %v to be deleted", permissionName)
		}
	}
}
