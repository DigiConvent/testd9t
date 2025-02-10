package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetUser(t *testing.T) {
	iamService := GetTestIAMService("iam")

	fakeUser := iam_domain.UserWrite{
		Email:       "TestGetUser@test.test",
		FirstName:   "Test",
		LastName:    "GetUser",
		DateOfBirth: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
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

	if user.Email != fakeUser.Email {
		t.Errorf("Email is not equal")
	}

	if user.FirstName != fakeUser.FirstName {
		t.Errorf("FirstName is not equal")
	}

	if user.LastName != fakeUser.LastName {
		t.Errorf("LastName is not equal")
	}

	if user.DateOfBirth.GoString() != fakeUser.DateOfBirth.GoString() {
		t.Errorf("DateOfBirth is not equal")
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
	user, status = iamService.GetUser(&adminId)

	if status.Err() {
		t.Errorf("Error: %v", status.Message)
	}

	if user == nil {
		t.Fatal("User is nil")
	}

	if user.Email != "" {
		t.Errorf("Email is not empty")
	}

	if user.FirstName != "" {
		t.Errorf("FirstName is not empty")
	}

	if user.LastName != "" {
		t.Errorf("LastName is not empty")
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

	if profile.Permissions[0].Name != "super" {
		t.Errorf("Permission is not super")
	}
}
