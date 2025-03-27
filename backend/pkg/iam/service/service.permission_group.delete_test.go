package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestDeletePermissionGroup(t *testing.T) {
	testService := GetTestIAMService("iam")

	res, _ := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroupDelete",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
		Permissions: []string{},
	})

	if res == nil {
		t.Fatal("Expected a result")
	}

	status := testService.DeletePermissionGroup(res)

	if status.Err() {
		t.Fatal(status.Message)
	}

	_, status = testService.GetPermissionGroup(res)

	if !status.Err() {
		t.Fatal("Expected an error, instead I got ", status.Code)
	}

	if status.Code != 404 {
		t.Fatal("Expected 404")
	}
}
