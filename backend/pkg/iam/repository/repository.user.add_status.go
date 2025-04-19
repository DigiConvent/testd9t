package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) AddUserStatusToUser(d *iam_domain.UserBecameStatusWrite) core.Status {
	_, err := r.db.Exec(`insert into user_became_status ("user", status, start, description) values (?, ?, ?, ?)`, d.User, d.UserStatus, d.Start, d.Description)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
