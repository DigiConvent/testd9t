package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetUserRoleProfile(t *testing.T) {
	testService := GetTestIAMService("iam")

	rootId := uuid.MustParse(getRootPermissionGroup())
	userRoleId, _ := testService.CreateUserRole(&iam_domain.UserRoleWrite{
		Name:   "UserRoleProfile",
		Abbr:   "UR",
		Parent: &rootId,
	})

	_, status := testService.GetUserRole(&uuid.Max)

	if !status.Err() {
		t.Fatal("Expected an error")
	}

	res, status := testService.GetUserRoleProfile(userRoleId)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}

	if len(res.History) != 0 {
		t.Fatal("Expected no history")
	}

	if res.UserRole.Name != "UserRoleProfile" {
		t.Fatal("Expected UserRoleProfile, instead I got ", res.UserRole.Name)
	}

	if res.UserRole.Abbr != "UR" {
		t.Fatal("Expected UR, instead I got ", res.UserRole.Abbr)
	}

	if res.UserRole.Description != "" {
		t.Fatal("Expected '', instead I got ", res.UserRole.Description)
	}

	if res.PermissionGroup.Parent == nil {
		t.Fatal("Expected a parent")
	}
}
