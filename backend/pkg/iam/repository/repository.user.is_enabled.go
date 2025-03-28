package iam_repository

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

var JwtDuration time.Duration = time.Hour * 3

type disabledUser struct {
	ID    uuid.UUID `json:"id"`
	since time.Time
}

var disabledUsers map[string]disabledUser
var lastCheck = time.Now()

func (r *IAMRepository) IsEnabled(id *uuid.UUID) (bool, core.Status) {
	if disabledUsers == nil {
		disabledUsers = make(map[string]disabledUser)
		result, err := r.db.Query(`select id from users where enabled = false`)
		if err != nil {
			return false, *core.InternalError(err.Error())
		}
		defer result.Close()
		for result.Next() {
			var user disabledUser
			err := result.Scan(&user.ID)
			user.since = time.Now()
			if err != nil {
				return false, *core.InternalError(err.Error())
			}
			disabledUsers[user.ID.String()] = user
		}
	}

	if id == nil {
		return false, *core.UnprocessableContentError("ID is required")
	}
	_, ok := disabledUsers[id.String()]

	// just run this every 3 hours
	if lastCheck.Add(JwtDuration).Before(time.Now()) {
		removeExpiredDisabledUsers()
	}
	return !ok, *core.StatusSuccess()
}

// no need to check since the list has to be initialised after the very first login
func disableUser(id *uuid.UUID) {
	disabledUsers[id.String()] = disabledUser{ID: *id, since: time.Now()}
}

func enableUser(id *uuid.UUID) {
	delete(disabledUsers, id.String())
}

func removeExpiredDisabledUsers() {
	for k, v := range disabledUsers {
		if time.Since(v.since) > JwtDuration {
			delete(disabledUsers, k)
		}
	}
	lastCheck = time.Now()
}
