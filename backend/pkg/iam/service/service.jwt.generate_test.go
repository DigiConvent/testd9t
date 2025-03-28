package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestGenerateJwt(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser := &iam_domain.UserWrite{
		Emailaddress: "GenerateJwt@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := iamService.CreateUser(testUser)

	// this will fail if the user is not enabled
	_, status := iamService.GenerateJwt(id)

	if !status.Err() {
		t.Fatal(status.Message)
	}

	iamService.SetEnabled(id, true)

	token, _ := iamService.GenerateJwt(id)

	if token == "" {
		t.Fatal("Expected a token")
	}

	_, status = iamService.GenerateJwt(nil)
	if !status.Err() {
		t.Fatal("Expected an error")
	}
}
