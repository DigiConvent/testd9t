package iam_service_test

import (
	"fmt"
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestCreatePermissionGroup(t *testing.T) {
	fmt.Println("TestCreatePermissionGroup")
	testService := GetTestIAMService("iam")

	res, status := testService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "PermissionGroup",
		Abbr:        "PG",
		Description: "test",
		IsGroup:     true,
		IsNode:      false,
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}

	fmt.Println(res)
}
