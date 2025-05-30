package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestUpdateUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	user := &iam_domain.UserWrite{
		Emailaddress: "TestUpdateUser@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}
	userId, _ := iamService.CreateUser(user)

	user.Emailaddress = "TestUpdateUser@test.test1"
	user.FirstName = "Updated"
	// user.LastName = "McUpdated2"

	status := iamService.UpdateUser(userId, user)
	if status.Err() {
		t.Fatal("Error updating user", status.Message)
	}

	updatedUser, _ := iamService.GetUser(userId)
	if updatedUser.FirstName != user.FirstName ||
		updatedUser.LastName != user.LastName ||
		updatedUser.Emailaddress != user.Emailaddress {
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
