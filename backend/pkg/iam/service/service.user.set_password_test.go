package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestSetUserPassword(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser, _ := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "TestSetUserPassword@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
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
