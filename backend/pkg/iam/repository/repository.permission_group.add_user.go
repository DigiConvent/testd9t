package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) core.Status {
	if permissionGroup == nil || userId == nil {
		return *core.UnprocessableContentError("permission group and user id must be provided")
	}

	var generated bool
	err := r.db.QueryRow(`select "generated" from permission_groups where id = ?`, permissionGroup.String()).Scan(&generated)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	if generated {
		return *core.UnprocessableContentError("cannot add user to generated permission groups")
	}

	_, err = r.db.Exec(`insert into permission_group_has_user (permission_group, user) values (?, ?)`, permissionGroup.String(), userId.String())

	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
