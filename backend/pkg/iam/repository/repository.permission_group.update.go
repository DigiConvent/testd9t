package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) core.Status {
	pg, _ := r.GetPermissionGroup(id)
	if pg.Meta != "" {
		arg.Name = pg.Name
		arg.Abbr = pg.Abbr
	}
	result, err := r.db.Exec(`update permission_groups set name = ?, abbr = ?, description = ? where id = ?`, arg.Name, arg.Abbr, arg.Description, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("permission group not found")
	}

	return *core.StatusNoContent()
}
