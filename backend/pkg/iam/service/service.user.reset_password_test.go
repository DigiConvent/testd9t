package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestResetPassword(t *testing.T) {
	iamService := GetTestIAMService("iam")

	uid, status := iamService.CreateUser(&iam_domain.UserWrite{
		Email:       "TestResetPassword@test.test",
		FirstName:   "Test",
		LastName:    "McTest",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	if status.Err() {
		t.Log(uid)
		t.Fatal(status.Message)
	}

	user, status := iamService.GetUser(uid)
	if status.Err() {
		t.Fatal(status.Message)
	}

	t.Log(user.Email)

	token, status := iamService.ResetPassword("TestResetPassword@test.test")
	if status.Err() {
		t.Fatal(status.Message)
	}

	if token == "" {
		t.Fatal("Expected a token")
	}
}
