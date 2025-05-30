package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestSetEnabled(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Emailaddress: "SetUserEnabled@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.CreateUser(user)

	userRead, _ := iamService.GetUser(userId)
	if userRead.Enabled {
		t.Errorf("User should initially be disabled")
	}

	iamService.IsEnabled(userId)

	status := iamService.SetEnabled(userId, true)
	if status.Err() {
		t.Errorf("Error enabling user: %v", status)
	}

	userRead, _ = iamService.GetUser(userId)
	if !userRead.Enabled {
		t.Errorf("User should be enabled, instead got: %v", userRead.Enabled)
	}

	status = iamService.SetEnabled(userId, false)
	if status.Err() {
		t.Errorf("Error disabling user: %v", status)
	}

	userRead, _ = iamService.GetUser(userId)
	if userRead.Enabled {
		t.Errorf("User should be disabled")
	}

	status = iamService.SetEnabled(nil, true)
	if !status.Err() {
		t.Errorf("Expected error when setting enabled for nil user")
	}
}
