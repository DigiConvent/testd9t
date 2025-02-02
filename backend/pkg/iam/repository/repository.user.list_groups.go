package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserGroups(userId *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status) {
	rows, err := r.DB.Query(`select
	id,
	name,
	abbr,
	is_group,
	implied,
	uhpg.parent
	FROM user_has_permission_groups uhpg
	RIGHT JOIN permission_groups pg ON pg.id = uhpg.permission_group 
	WHERE uhpg.user = ?`, userId.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	var permissionGroups = make([]*iam_domain.PermissionGroupFacade, 0)
	for rows.Next() {
		var permissionGroup iam_domain.PermissionGroupFacade
		var parentID uuid.UUID
		err := rows.Scan(&permissionGroup.ID, &permissionGroup.Name, &permissionGroup.Abbr, &permissionGroup.IsGroup, &permissionGroup.Implied, &parentID)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		if parentID != uuid.Nil {
			permissionGroup.Parent = &parentID
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, *core.StatusSuccess()
}
