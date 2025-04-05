package iam_service_test

import "testing"

func TestListUserRolees(t *testing.T) {
	iamService := GetTestIAMService("iam")

	UserRoleses, status := iamService.ListUserRoles()

	if status.Err() {
		t.Errorf("ListUserRoleses() failed: %s", status.Message)
	}

	if len(UserRoleses) == 0 {
		t.Errorf("ListUserRoleses() failed: no user statusses found")
	}
}
