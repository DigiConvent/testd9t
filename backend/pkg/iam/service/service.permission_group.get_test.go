package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetPermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	res, status := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupGet",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
		Parent:      getRootPermissionGroup(),
	})

	if res == nil {
		t.Log(status.Message)
		t.Fatal("Expected a result")
	}

	permissionGroup, status := testService.GetPermissionGroup(res)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if permissionGroup.Name != "PermissionGroupGet" {
		t.Fatal("Expected PermissionGroupGet, instead I got ", permissionGroup.Name)
	}

	if permissionGroup.Abbr != "PG" {
		t.Fatal("Expected PG, instead I got ", permissionGroup.Abbr)
	}

	if permissionGroup.Description != "test" {
		t.Fatal("Expected test, instead I got ", permissionGroup.Description)
	}

	randomFailingId, _ := uuid.NewV7()
	permissionGroup, status = testService.GetPermissionGroup(&randomFailingId)

	if !status.Err() {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	if status.Code != 404 {
		t.Fatal("Expected 404")
	}

	if permissionGroup != nil {
		t.Fatal("Expected nil, instead I got ", permissionGroup)
	}
}
