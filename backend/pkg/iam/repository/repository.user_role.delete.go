package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) DeleteUserRole(id *uuid.UUID) core.Status {
	result, err := r.db.Exec(`delete from permission_groups where id = ? and meta = 'role'`, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("No such user role")
	}

	return core.Status{Code: 204, Message: "OK"}
}
