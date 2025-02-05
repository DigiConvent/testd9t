package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestCreateUserStatus(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUserStatus := &iam_domain.UserStatusWrite{
		Name:        "UserStatusCreate",
		Abbr:        "USC",
		Description: "test",
		Archived:    false,
	}

	id, status := iamService.CreateUserStatus(testUserStatus)

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

	if pg.Name != testUserStatus.Name {
		t.Errorf("Expected %s, got %s", testUserStatus.Name, pg.Name)
	}
}
