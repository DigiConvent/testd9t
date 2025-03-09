package iam_service_test

import "testing"

func TestListUserStatuses(t *testing.T) {
	iamService := GetTestIAMService("iam")

	userStatusses, status := iamService.ListUserStatuses()

	if status.Err() {
		t.Errorf("ListUserStatusses() failed: %s", status.Message)
	}

	if len(userStatusses) == 0 {
		t.Errorf("ListUserStatusses() failed: no user statusses found")
	}
}
