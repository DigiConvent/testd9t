package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetUserProfile(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Emailaddress: "GetUserProfile@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.CreateUser(user)

	profile, err := iamService.GetUserProfile(userId)
	if err.Err() {
		t.Fatalf("GetUserProfile failed: %v", err)
	}

	if profile == nil {
		t.Fatalf("GetUserProfile failed: profile is nil")
	}

	if profile.User == nil {
		t.Fatalf("GetUserProfile failed: user is nil")
	}
}
