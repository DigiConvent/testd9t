package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) core.Status {
	if arg.Parent == nil || arg.Parent.String() == uuid.Nil.String() {
		arg.Parent = nil
	}
	if arg.Id == nil {
		return *core.UnprocessableContentError("permission group ID is required")
	}

	result, err := r.db.Exec("update permission_groups set parent = ? where id = ? and exists (select 1 from permission_groups where id = ? and meta is null)", arg.Parent, arg.Id, arg.Parent)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("permission group " + arg.Id.String() + " not found")
	}
	return *core.StatusNoContent()
}
