package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UpdateUserRole(id *uuid.UUID, arg *iam_domain.UserRoleWrite) core.Status {
	result, err := r.db.Exec("update permission_groups set name = ?, abbr = ?, description = ? where id = ? and meta = 'role'", arg.Name, arg.Abbr, arg.Description, id.String())
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("User role not found")
	}
	return *core.StatusNoContent()
}
