package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetUserProfile(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Email:       "GetUserProfile@test.test",
		FirstName:   "Test",
		LastName:    "McTest",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
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
