package iam_service_test

import (
	"testing"
)

func TestListPermissions(t *testing.T) {
	iamService := GetTestIAMService("iam")

	permissions, status := iamService.ListPermissions()

	if status.Err() {
		t.Errorf("ListPermissions() failed: %s", status.Message)
	}

	if len(permissions) == 0 {
		t.Errorf("ListPermissions() failed: no permissions found")
	}
}
