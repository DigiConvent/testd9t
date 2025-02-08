package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestUpdateUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Email:       "TestUpdateUser@test.test",
		FirstName:   "Test",
		LastName:    "McTest",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	userId, _ := iamService.CreateUser(user)

	user.Email = "TestUpdateUser@test.test1"
	user.FirstName = "Updated"
	user.LastName = "McUpdated"
	user.DateOfBirth = time.Date(2001, 2, 2, 0, 0, 0, 0, time.UTC)

	status := iamService.UpdateUser(userId, user)
	if status.Err() {
		t.Fatal("Error updating user", status.Message)
	}

	updatedUser, _ := iamService.GetUser(userId)
	if updatedUser.FirstName != user.FirstName ||
		updatedUser.LastName != user.LastName ||
		updatedUser.Email != user.Email ||
		updatedUser.DateOfBirth.String() != time.Date(2001, 2, 2, 0, 0, 0, 0, time.UTC).String() {
		t.Fatal("User not updated")
	}

	status = iamService.UpdateUser(nil, user)
	if !status.Err() {
		t.Fatal("Should have errored")
	}

	status = iamService.UpdateUser(userId, nil)
	if !status.Err() {
		t.Fatal("Should have errored")
	}
}
