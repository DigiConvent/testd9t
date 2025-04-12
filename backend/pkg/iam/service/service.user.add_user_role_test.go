package iam_service_test

import (
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestUserAddUserRole(t *testing.T) {
	iamService := GetTestIAMService("iam")

	testUser := &iam_domain.UserWrite{
		Emailaddress: "UserAddUserRole@test.test",
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

	parent := uuid.MustParse(getRootPermissionGroup())
	testCurrentUserRole := &iam_domain.UserRoleWrite{
		Name:        "UserAddUserRoleCurrent",
		Abbr:        "UAUS",
		Description: "test",
		Parent:      &parent,
	}

	currentUserRoleId, status := iamService.CreateUserRole(testCurrentUserRole)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if currentUserRoleId == nil {
		t.Fatal("Expected a result")
	}

	status = iamService.AddUserRole(&iam_domain.AddUserRoleToUser{
		UserID: *id,
		RoleID: *currentUserRoleId,
		Start:  time.Now().Add(-2 * time.Hour),
		End:    time.Now().Add(2 * time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	// attempt to add the same role again but with overlapping from the previous one
	status = iamService.AddUserRole(&iam_domain.AddUserRoleToUser{
		UserID: *id,
		RoleID: *currentUserRoleId,
		Start:  time.Now(),
		End:    time.Now().Add(4 * time.Hour),
	})

	if !status.Err() {
		t.Fatal("expected an error, instead got", status.Code)
	}

	status = iamService.AddUserRole(&iam_domain.AddUserRoleToUser{
		UserID: *id,
		RoleID: *currentUserRoleId,
		Start:  time.Now().Add(-3 * time.Hour),
		End:    time.Now(),
	})

	if !status.Err() {
		t.Fatal("expected an error")
	}

	testFutureUserRole := &iam_domain.UserRoleWrite{
		Name:        "UserAddUserRoleFuture",
		Abbr:        "UAUSF",
		Description: "test",
		Parent:      &parent,
	}

	futureUserRoleId, _ := iamService.CreateUserRole(testFutureUserRole)

	if futureUserRoleId == nil {
		t.Fatal("Expected a result")
	}

	status = iamService.AddUserRole(&iam_domain.AddUserRoleToUser{
		UserID: *id,
		RoleID: *futureUserRoleId,
		Start:  time.Now().Add(5 * time.Hour),
	})

	if status.Err() {
		t.Fatal(status.Message)
	}

	futurePG, status := iamService.GetPermissionGroupProfile(futureUserRoleId)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if futurePG == nil {
		t.Fatal("Expected a result")
	}

	if len(futurePG.Members) != 0 {
		t.Fatalf("Expected 0 user, got %d", len(futurePG.Members))
	}

	currentPG, status := iamService.GetPermissionGroupProfile(currentUserRoleId)

	if status.Err() {
		t.Fatal(status.Message)
	}

	if currentPG == nil {
		t.Fatal("Expected a result")
	}

	if len(currentPG.Members) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(currentPG.Members))
	}

	if status.Err() {
		t.Fatal(status.Message)
	}
}
