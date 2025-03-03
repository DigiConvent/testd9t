package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestSetUserPassword(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser, _ := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "TestSetUserPassword@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
		DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	status := iamService.SetUserPassword(testUser, "password123")

	if status.Err() {
		t.Fatal(status.Message)
	}

	status = iamService.SetUserPassword(nil, "password123")

	if !status.Err() {
		t.Fatal("Expected an error")
	}
}
