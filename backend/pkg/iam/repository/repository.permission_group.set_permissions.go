package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) SetPermissionsForPermissionGroup(permissionGroupId *uuid.UUID, permissions []*uuid.UUID) core.Status {
	result, err := r.DB.Exec(`
WITH existing_permissions AS (
    SELECT p.id
    FROM permissions p
    JOIN permission_group_has_permission pghp ON p.id = pgp.permission
    WHERE permission_group = $1
),
new_permissions AS (
    SELECT unnest($2::varchar[]) AS permission
),
to_delete AS (
    SELECT permission
    FROM existing_permissions
    EXCEPT
    SELECT permission
    FROM new_permissions
),
to_add AS (
    SELECT permission
    FROM new_permissions
    EXCEPT
    SELECT permission
    FROM existing_permissions
)

DELETE FROM permission_group_has_permission
WHERE permission_group = $1 AND permission IN (SELECT permission FROM to_delete);

INSERT INTO permission_group_has_permission (permission_group, permission)
SELECT $1, permission
FROM to_add;`, permissionGroupId.String(), permissions)

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("Permission group not found")
	}
	return *core.StatusSuccess()
}
