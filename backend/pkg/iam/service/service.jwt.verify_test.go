package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestVerifyJwt(t *testing.T) {
	iamService := GetTestIAMService("iam")

	_, status := iamService.VerifyJwt("badtoken")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	_, status = iamService.VerifyJwt("")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	testUser := &iam_domain.UserWrite{
		Emailaddress: "VerifyJwt@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	id, _ := iamService.CreateUser(testUser)
	iamService.SetEnabled(id, true)
	token, _ := iamService.GenerateJwt(id)
	if !status.Err() {
		t.Fatal("this should fail because the user is not enabled")
	}

	theId, status := iamService.VerifyJwt(token)
	if status.Err() {
		t.Fatal(status.Message)
	}

	if theId == nil {
		t.Fatal("Expected a result")
	}

	if theId.String() != id.String() {
		t.Fatal("Expected ", id.String(), " instead I got ", theId.String())
	}
}
