package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestLoginUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser, _ := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "TestLoginUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
		DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	iamService.SetUserPassword(testUser, "password123")

	uid, status := iamService.LoginUser("TestLoginUser@test.test", "password123")

	if status.Err() {
		t.Fatal(status.Message)
	}

	if uid == nil {
		t.Fatal("Expected a result")
	}

	if uid.String() != testUser.String() {
		t.Fatal("Expected the same ID")
	}
}
