package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestCreateUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUserRole := &iam_domain.UserRoleWrite{
		PermissionGroupWrite: iam_domain.PermissionGroupWrite{
			Name:        "UserRoleCreate",
			Abbr:        "USC",
			Description: "test",
			Parent:      getRootPermissionGroupUuid(),
		},
	}

	id, status := iamService.CreateUserRole(testUserRole)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if id == nil {
		t.Fatal("Expected a result")
	}

	pg, status := iamService.GetPermissionGroup(id)
	if status.Err() {
		t.Fatal(status.Message)
	}

	if pg == nil {
		t.Fatal("Expected a result")
	}

	if pg.Name != testUserRole.Name {
		t.Errorf("Expected %s, got %s", testUserRole.Name, pg.Name)
	}
}
