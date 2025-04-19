package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestResetPassword(t *testing.T) {
	iamService := GetTestIAMService("iam")

	uid, status := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "TestResetPassword@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	_, status = iamService.GetUser(uid)
	if status.Err() {
		t.Fatal(status.Message)
	}

	token, status := iamService.ResetPassword("TestResetPassword@test.test")
	if status.Err() {
		t.Fatal(status.Message)
	}

	if token == "" {
		t.Fatal("Expected a token")
	}
}
