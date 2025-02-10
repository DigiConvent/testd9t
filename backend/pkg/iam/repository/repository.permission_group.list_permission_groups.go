package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListPermissionGroupPermissionGroups(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status) {
	var permissionGroups = make([]*iam_domain.PermissionGroupFacade, 0)
	rows, err := r.db.Query(`SELECT id, name, parent, implied FROM permission_group_has_permission_groups where child_id = ?`, arg.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var permissionGroup iam_domain.PermissionGroupFacade
		err := rows.Scan(&permissionGroup.ID, &permissionGroup.Name, &permissionGroup.Parent, &permissionGroup.Implied)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, *core.StatusSuccess()
}
