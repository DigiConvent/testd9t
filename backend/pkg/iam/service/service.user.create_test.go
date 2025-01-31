package iam_service_test

import (
	"fmt"
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

var testUser = &iam_domain.UserWrite{
	FirstName:   "FirstName",
	LastName:    "LastName",
	Email:       "a@a.a",
	DateOfBirth: time.Now(),
}

func TestCreateUser(t *testing.T) {
	fmt.Println("TestCreateUser")
	testService := GetTestIAMService("iam")

	res, status := testService.CreateUser(testUser)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if res == nil {
		t.Fatal("Expected a result")
	}

	t.Log(res)
}
