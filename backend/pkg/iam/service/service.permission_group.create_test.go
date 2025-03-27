package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestCreatePermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	testPermissionGroup := &iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupCreate",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
		Parent:      getRootPermissionGroup(),
	}

	res, status := testService.CreatePermissionGroup(testPermissionGroup)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}
}
