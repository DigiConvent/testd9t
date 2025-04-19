package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionPermissionGroups(name string) ([]*iam_domain.PermissionGroupFacade, core.Status) {
	rows, err := r.db.Query(`select pg.id, pg.name, pg.abbr, pg.meta, pg.parent, pg.generated from permission_group_has_permission pghp join permission_groups pg on pg.id = pghp.permission_group where pghp.permission = ?`, name)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	permissionGroups := make([]*iam_domain.PermissionGroupFacade, 0)
	for rows.Next() {
		var permissionGroup iam_domain.PermissionGroupFacade
		err := rows.Scan(&permissionGroup.Id, &permissionGroup.Name, &permissionGroup.Abbr, &permissionGroup.Meta, &permissionGroup.Parent, &permissionGroup.Generated)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}

	return permissionGroups, *core.StatusSuccess()
}
