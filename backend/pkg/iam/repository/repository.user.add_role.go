package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) AddUserRoleToUser(d *iam_domain.AddUserRoleToUser) core.Status {
	_, err := r.db.Exec(`insert into user_became_role ("user", "role", start) values (?, ?, ?)`, d.UserID, d.RoleID, d.When)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
