package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestRegisterTelegramUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Emailaddress: "TestRegisterTelegramUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
		DateOfBirth:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	userId, _ := iamService.CreateUser(user)

	code, _ := iamService.GetTelegramRegistrationCode(userId)

	status := iamService.RegisterTelegramUser(15, user.Emailaddress, code)
	if status.Err() {
		t.Fatal("Error registering user", status.Message)
	}

}
