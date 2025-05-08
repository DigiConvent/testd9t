package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) core.Status {
	if permissionGroup == nil || userId == nil {
		return *core.UnprocessableContentError("permission group and user id must be provided")
	}

	// make sure that the permission group has a valid meta
	pg, status := r.GetPermissionGroup(permissionGroup)
	if status.Err() {
		return status
	}

	if pg.Meta == "" {
		return *core.UnprocessableContentError("iam.permission_group.add_user.invalid_meta")
	}

	_, err := r.db.Exec(`insert into permission_group_has_user (permission_group, user) values (?, ?)`, permissionGroup.String(), userId.String())

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
