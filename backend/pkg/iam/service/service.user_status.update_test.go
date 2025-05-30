package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestUpdateUserStatus(t *testing.T) {
	iamService := GetTestIAMService("iam")

	rootId := uuid.MustParse(getRootPermissionGroup())
	testUserStatus := &iam_domain.UserStatusWrite{
		PermissionGroupWrite: iam_domain.PermissionGroupWrite{
			Name:        "UserStatusUpdate",
			Abbr:        "USU",
			Description: "test",
			Parent:      &rootId,
		},
		Archived: false,
	}

	id, _ := iamService.CreateUserStatus(testUserStatus)

	if id == nil {
		t.Fatal("Expected a result")
	}

	testUserStatus.Name = "UserStatusUpdateUpdated"
	testUserStatus.Abbr = "USUU"
	testUserStatus.Description = "testx"
	testUserStatus.Archived = true

	fakeId := uuid.Max
	status := iamService.UpdateUserStatus(&fakeId, testUserStatus)

	if !status.Err() {
		t.Fatal("Expected an error")
	}

	status = iamService.UpdateUserStatus(id, testUserStatus)

	if status.Err() {
		t.Fatal(status.Message)
	}

	userStatus, _ := iamService.GetUserStatus(id)

	if userStatus == nil {
		t.Errorf("Expected a result, got %v", userStatus)
	}

	if userStatus.Id.String() != id.String() {
		t.Errorf("Expected %v, got %v", id, userStatus.Id.String())
	}

	if userStatus.Name != "UserStatusUpdateUpdated" {
		t.Errorf("Expected %s, got %s", "UserStatusUpdateUpdated", userStatus.Name)
	}

	if userStatus.Abbr != "USUU" {
		t.Errorf("Expected %s, got %s", "USUU", userStatus.Abbr)
	}

	if userStatus.Description != "testx" {
		t.Errorf("Expected %s, got %s", "testx", userStatus.Description)
	}

	if userStatus.Archived != true {
		t.Errorf("Expected %t, got %t", true, userStatus.Archived)
	}

	pg, _ := iamService.GetPermissionGroup(id)

	if pg == nil {
		t.Fatal("Expected a result")
	}

	if pg.Name != testUserStatus.Name {
		t.Errorf("Expected %s, got %s", testUserStatus.Name, pg.Name)
	}

	if pg.Abbr != testUserStatus.Abbr {
		t.Errorf("Expected %s, got %s", testUserStatus.Abbr, pg.Abbr)
	}

	if pg.Description != testUserStatus.Description {
		t.Errorf("Expected %s, got %s", testUserStatus.Description, pg.Description)
	}
}
