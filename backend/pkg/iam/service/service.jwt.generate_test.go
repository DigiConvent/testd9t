package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGenerateJwt(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser := &iam_domain.UserWrite{
		Email:       "GenerateJwt@test.test",
		FirstName:   "Test",
		LastName:    "McTest",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	id, _ := iamService.CreateUser(testUser)

	token, status := iamService.GenerateJwt(id)
	if status.Err() {
		t.Fatal(status.Message)
	}

	if token == "" {
		t.Fatal("Expected a token")
	}

	_, status = iamService.GenerateJwt(nil)
	if !status.Err() {
		t.Fatal("Expected an error")
	}
}
