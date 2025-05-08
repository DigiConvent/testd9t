package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) AddRoleToUser(d *iam_domain.AddRoleToUserWrite) core.Status {
	_, err := r.db.Exec(`insert into permission_group_has_user (user, permission_group, start, "end") values (?, ?, ?, ?)`, d.User, d.Role, d.Start, d.End)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
