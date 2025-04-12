package iam_service_test

import (
	"testing"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	fakeUser := iam_domain.UserWrite{
		Emailaddress: "TestGetUser@test.test",
		FirstName:    "Test",
		LastName:     "GetUser",
	}

	id, _ := iamService.CreateUser(&fakeUser)

	user, status := iamService.GetUser(id)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if user == nil {
		t.Fatal("User is nil")
	}

	if user.ID != *id {
		t.Errorf("ID is not equal")
	}

	if user.Emailaddress != fakeUser.Emailaddress {
		t.Errorf("Email is not equal")
	}

	if user.FirstName != fakeUser.FirstName {
		t.Errorf("FirstName is not equal")
	}

	if user.LastName != fakeUser.LastName {
		t.Errorf("LastName is not equal")
	}

	// test fake uuid

	id = nil
	user, status = iamService.GetUser(id)

	if !status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if user != nil {
		t.Fatal("User is not nil")
	}

	unknownId, _ := uuid.NewV7()

	user, status = iamService.GetUser(&unknownId)

	if !status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if status.Code != 404 {
		t.Errorf("Status code is not 404")
	}

	if user != nil {
		t.Fatal("User is not nil")
	}

	// also test admin user
	adminId := uuid.Nil
	admin, status := iamService.GetUser(&adminId)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if admin == nil {
		t.Fatal("User is nil")
	}

	if admin.Emailaddress != "" {
		t.Errorf("Email is not empty")
	}

	profile, status := iamService.GetUserProfile(&adminId)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if profile == nil {
		t.Fatal("Profile is nil")
	}

	if profile.User == nil {
		t.Fatal("User is nil")
	}

	if profile.Permissions[0].Name != "admin" {
		t.Errorf("Permission is not admin")
	}
}
