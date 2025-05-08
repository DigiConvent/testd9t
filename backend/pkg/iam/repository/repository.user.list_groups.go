package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserGroups(userId *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status) {
	rows, err := r.db.Query(`select pg.id, pg.name, pg.abbr, pg.meta, uhpg.implied, uhpg.parent 
 from user_has_permission_groups uhpg 
 right join permission_groups pg on pg.id = uhpg.permission_group
 where uhpg.user = ?`, userId.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	var permissionGroups = make([]*iam_domain.PermissionGroupFacade, 0)
	for rows.Next() {
		var permissionGroup iam_domain.PermissionGroupFacade
		var parentID uuid.UUID
		err := rows.Scan(&permissionGroup.Id, &permissionGroup.Name, &permissionGroup.Abbr, &permissionGroup.Meta, &permissionGroup.Implied, &parentID)
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
