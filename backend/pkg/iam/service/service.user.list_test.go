package iam_service_test

import "testing"

func TestListUsers(t *testing.T) {
	iamService := GetTestIAMService("iam")

	userList, status := iamService.ListUsers(nil)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if userList == nil {
		t.Fatal("Expected a page of users, got nil")
	}

}
