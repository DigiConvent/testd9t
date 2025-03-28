package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func TestUserAddUserStatus(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser := &iam_domain.UserWrite{
		Emailaddress: "UserAddUserStatus@test.test",
		FirstName:    "Test",
		LastName:     "McTest",
	}

	id, status := iamService.CreateUser(testUser)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if id == nil {
		t.Fatal("Expected a result")
	}

	testCurrentUserStatus := &iam_domain.UserStatusWrite{
		Name:        "UserAddUserStatusCurrent",
		Abbr:        "UAUS",
		Description: "test",
		Archived:    false,
	}

	currentUserStatusId, status := iamService.CreateUserStatus(testCurrentUserStatus)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if currentUserStatusId == nil {
		t.Fatal("Expected a result")
	}

	status = iamService.AddUserStatus(&iam_domain.AddUserStatusToUser{
		UserID:   *id,
		StatusID: *currentUserStatusId,
		When:     time.Now().Add(3 * -time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	testFutureUserStatus := &iam_domain.UserStatusWrite{
		Name:        "UserAddUserStatusFuture",
		Abbr:        "UAUSF",
		Description: "test",
		Archived:    false,
	}

	futureUserStatusId, _ := iamService.CreateUserStatus(testFutureUserStatus)

	if futureUserStatusId == nil {
		t.Fatal("Expected a result")
	}

	status = iamService.AddUserStatus(&iam_domain.AddUserStatusToUser{
		UserID:   *id,
		StatusID: *futureUserStatusId,
		When:     time.Now().Add(3 * time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	futurePG, status := iamService.GetPermissionGroupProfile(futureUserStatusId)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if futurePG == nil {
		t.Fatal("Expected a result")
	}

	if len(futurePG.Members) != 0 {
		t.Fatalf("Expected 0 user, got %d", len(futurePG.Members))
	}

	currentPG, status := iamService.GetPermissionGroupProfile(currentUserStatusId)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if currentPG == nil {
		t.Fatal("Expected a result")
	}

	userProfile, status := iamService.GetUserProfile(id)

	if status.Err() {
		t.Fatal(status.Message)
	}
	for _, group := range userProfile.Groups {
		t.Log(group.Name == currentPG.PermissionGroup.Name)
	}

	if len(currentPG.Members) != 1 {
		t.Log(currentPG.Members)
		t.Fatalf("Expected 1 user, got %d", len(currentPG.Members))
	}

	if status.Err() {
		t.Fatal(status.Message)
	}

}
