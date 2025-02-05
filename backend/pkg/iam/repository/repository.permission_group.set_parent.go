package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) core.Status {
	if arg.Parent.String() == uuid.Nil.String() {
		arg.Parent = nil
	}
	result, err := r.DB.Exec("UPDATE permission_groups SET parent = ? WHERE id = ?", arg.Parent, arg.ID)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("permission group not found")
	}
	return *core.StatusNoContent()
}
