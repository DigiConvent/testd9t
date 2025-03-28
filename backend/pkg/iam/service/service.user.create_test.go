package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

var testUser = &iam_domain.UserWrite{
	FirstName:    "FirstName",
	LastName:     "LastName",
	Emailaddress: "a@a.a",
}

func TestCreateUser(t *testing.T) {
	testService := GetTestIAMService("iam")

	res, status := testService.CreateUser(testUser)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}
}
