package iam_service_test

import (
	"regexp"
	"strconv"
	"testing"
	"time"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func TestGetUserStatusProfile(t *testing.T) {
	rootId := uuid.MustParse(getRootPermissionGroup())
	userStatusId, _ := GetTestIAMService("iam").CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "Test User Status Profile",
		Abbr:        "TUS",
		Description: "testxs",
		Archived:    true,
		Parent:      &rootId,
	})
	userStatusId2, _ := GetTestIAMService("iam").CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "Test User Status Profile2",
		Abbr:        "TUS",
		Description: "testxs",
		Archived:    true,
		Parent:      &rootId,
	})
	userStatusId3, _ := GetTestIAMService("iam").CreateUserStatus(&iam_domain.UserStatusWrite{
		Name:        "Test User Status Profile3",
		Abbr:        "TUS",
		Description: "testxs",
		Archived:    true,
		Parent:      &rootId,
	})

	userStatusProfile, status := GetTestIAMService("iam").GetUserStatusProfile(userStatusId)

	if status.Err() {
		t.Fatalf("GetUserStatusProfile() failed: %s", status.Message)
	}

	if userStatusProfile == nil {
		t.Fatalf("GetUserStatusProfile() failed: profile is nil")
	}

	if userStatusProfile.UserStatus == nil {
		t.Fatalf("GetUserStatusProfile() failed: user status is nil")
	}

	// check if the users are present
	if len(userStatusProfile.History) != 0 {
		t.Fatalf("expected 0 users in its history, instead got %v", len(userStatusProfile.History))
	}

	// add 3 users...
	userIds := make([]uuid.UUID, 4)
	for i := 0; i < 4; i++ {
		userId, _ := GetTestIAMService("iam").CreateUser(&iam_domain.UserWrite{
			Emailaddress: "GetUserStatusProfile" + strconv.Itoa(i) + "@test.test",
			FirstName:    "Test",
			LastName:     "McTest",
		})

		userIds[i] = *userId

		// link them all to the first user status
		GetTestIAMService("iam").AddUserBecameStatus(&iam_domain.UserBecameStatusWrite{
			UserStatus: *userStatusId,
			User:       userIds[i],
			Start:      time.Now().Add(-1 * time.Hour),
		})
	}

	// add a user status as the later one. Since this is a more recent one, this user should not show up in the profile of the first user status
	GetTestIAMService("iam").AddUserBecameStatus(&iam_domain.UserBecameStatusWrite{
		UserStatus:  *userStatusId2,
		User:        userIds[2],
		Start:       time.Now().Add(-30 * time.Minute),
		Description: "This should be in the past and be the most current user status for " + formatUuuid(userIds[2].String()),
	})
	// and since this user status is in the future, this user should show up
	GetTestIAMService("iam").AddUserBecameStatus(&iam_domain.UserBecameStatusWrite{
		UserStatus:  *userStatusId3,
		User:        userIds[3],
		Start:       time.Now().Add(30 * time.Minute),
		Description: "This should be in the future",
	})

	profile, status := GetTestIAMService("iam").GetUserStatusProfile(userStatusId)

	for _, userId := range userIds {
		t.Log(formatUuuid(userId.String()))
	}

	if status.Err() {
		t.Fatalf("GetUserStatusProfile() failed: %s", status.Message)
	}

	if len(profile.History) != 3 {
		t.Fatalf("expected 3 users, instead got %v", len(profile.History))
	}
}

func formatUuuid(s any) string {
	uuidRegex := `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`
	st, ok := s.(string)
	if !ok {
		return s.(string)
	}
	matched, err := regexp.MatchString(uuidRegex, st)
	if err != nil || !matched {
		return s.(string)
	}
	return st[:4] + "..." + st[len(st)-4:]
}
