package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) DeletePermissionGroup(arg *uuid.UUID) core.Status {
	res, err := r.DB.Exec(`DELETE FROM permission_groups WHERE id = ? AND "generated" = false`, arg.String())
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("Permission group not found")
	}
	return *core.StatusNoContent()
}
