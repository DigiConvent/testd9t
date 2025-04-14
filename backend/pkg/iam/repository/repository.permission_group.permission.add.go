package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) AddPermissionToPermissionGroup(permissionGroupId *uuid.UUID, permission string) core.Status {
	// insert the initial permission
	res, err := r.db.Exec(`insert into permission_group_has_permission (permission_group, permission) values (?, ?)`, permissionGroupId.String(), permission)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	d, err := res.RowsAffected()
	if err != nil || d == 0 {
		return *core.InternalError(err.Error())
	}

	// remove all the permissions that are implied by this permission
	_, err = r.db.Exec(`delete from permission_group_has_permission where permission_group = ? and permission like ?`, permissionGroupId.String(), permission+".%")

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusNoContent()
}
