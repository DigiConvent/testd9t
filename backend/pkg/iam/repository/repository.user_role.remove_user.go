package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) RemoveUserFromUserRole(arg *iam_domain.UserBecameRoleWrite) core.Status {
	result, err := r.db.Exec(`delete from user_became_role where "user" = ? and "role" = ? and start = ? and end = ?`, arg.User, arg.Role, arg.Start, arg.End)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return *core.InternalError(err.Error())
	}

	if rowsAffected == 0 {

		return *core.NotFoundError("no such combination of user and role")
	}
	return *core.StatusSuccess()
}
