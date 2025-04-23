package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestRemoveUserFromUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser, _ := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "TestRemoveUserFromUserRole@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	if testUser == nil {
		t.Fatal("Expected a result")
	}

	testRole, _ := iamService.CreateUserRole(&iam_domain.UserRoleWrite{
		Name:        "TestRemoveUserFromUserRole",
		Description: "TestRemoveUserFromUserRole",
		Abbr:        "TRU",
		Parent:      getRootPermissionGroupUuid(),
	})

	if testRole == nil {
		t.Fatal("Expected a result")
	}
}
