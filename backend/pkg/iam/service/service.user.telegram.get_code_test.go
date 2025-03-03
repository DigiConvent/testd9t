package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGetTelegramRegistrationCode(t *testing.T) {
	iamService := GetTestIAMService("iam")

	userId, _ := iamService.CreateUser(&iam_domain.UserWrite{
		Emailaddress: "GetTelegramRegistrationCode@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
		DateOfBirth:  time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	code, _ := iamService.GetTelegramRegistrationCode(userId)
	if code == "" {
		t.Fatal("Code is empty")
	}

	_, status := iamService.GetTelegramRegistrationCode(nil)
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	// no need to test the working of the code, I know it works
}
