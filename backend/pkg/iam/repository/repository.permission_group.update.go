package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) core.Status {
	result, err := r.db.Exec(`update permission_groups set name = ?, abbr = ?, description = ? where id = ? and generated = false`, arg.Name, arg.Abbr, arg.Description, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("Permission group not found")
	}

	return *core.StatusNoContent()
}
