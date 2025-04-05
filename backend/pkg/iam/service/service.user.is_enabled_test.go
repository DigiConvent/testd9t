package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestIsEnabled(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Emailaddress: "UserIsEnabled@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.CreateUser(user)

	userRead, _ := iamService.GetUser(userId)
	if userRead.Enabled {
		t.Errorf("User should initially be disabled")
	}

	isEnabled, _ := iamService.IsEnabled(userId)
	if isEnabled {
		t.Errorf("User should be disabled initially")
	}

	status := iamService.SetEnabled(userId, true)
	if status.Err() {
		t.Errorf("Error enabling user: %v", status)
	}

	isEnabled, _ = iamService.IsEnabled(userId)
	if !isEnabled {
		t.Errorf("User should be enabled")
	}

	status = iamService.SetEnabled(userId, false)
	if status.Err() {
		t.Errorf("Error disabling user: %v", status)
	}

	isEnabled, _ = iamService.IsEnabled(userId)
	if isEnabled {
		t.Errorf("User should be disabled")
	}

	status = iamService.SetEnabled(nil, true)
	if !status.Err() {
		t.Errorf("Expected error when setting enabled for nil user")
	}

	isEnabled, status = iamService.IsEnabled(nil)
	if !status.Err() {
		t.Errorf("Expected error when getting enabled for nil user")
	}
	if isEnabled {
		t.Errorf("Expected false when getting enabled for nil user")
	}
}
