package iam_repository

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

type disabledUser struct {
	ID    uuid.UUID `json:"id"`
	since time.Time
}

var disabledUsers map[string]disabledUser

func (r *IAMRepository) IsEnabled(id *uuid.UUID) (bool, core.Status) {
	if disabledUsers == nil {
		disabledUsers = make(map[string]disabledUser)
		result, err := r.db.Query(`select * from users where enabled = false`)
		if err != nil {
			return false, *core.InternalError(err.Error())
		}
		defer result.Close()
		for result.Next() {
			var user disabledUser
			err := result.Scan(&user.ID, &user.since)
			if err != nil {
				return false, *core.InternalError(err.Error())
			}
			disabledUsers[user.ID.String()] = user
		}
	}
	panic("unimplemented")
}

func disableUser(id *uuid.UUID) {
	panic("unimplemented")
}
