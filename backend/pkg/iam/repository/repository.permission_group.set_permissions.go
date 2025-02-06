package iam_repository

import (
	"encoding/json"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) SetPermissionsForPermissionGroup(permissionGroupId *uuid.UUID, permissions []string) core.Status {
	if permissionGroupId == nil {
		return *core.UnprocessableContentError("permissionGroupId cannot be nil")
	}

	setPermissions, _ := json.Marshal(permissions)
	pgId := permissionGroupId.String()
	_, err := r.DB.Exec(`
with existing_permissions as (select permission from permission_group_has_permission where permission_group = ?),
new_permissions as (select value as permission from json_each(?)),

to_delete as (select permission from existing_permissions except select permission from new_permissions),
to_add as (select permission from new_permissions except select permission from existing_permissions)

delete from permission_group_has_permission
where permission_group = ? and permission in (select permission from to_delete);`, pgId, string(setPermissions), pgId)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	if len(permissions) > 0 {
		_, err = r.DB.Exec(`
        with existing_permissions as (select permission from permission_group_has_permission where permission_group = ?),
        new_permissions as (select value as permission from json_each(?)),
        
        to_delete as (select permission from existing_permissions except select permission from new_permissions),
        to_add as (select permission from new_permissions except select permission from existing_permissions)
        
        insert into permission_group_has_permission (permission_group, permission)
        select ?, permission from to_add;`, pgId, string(setPermissions), pgId)

		if err != nil {
			return *core.InternalError(err.Error())
		}
	}
	return *core.StatusSuccess()
}
