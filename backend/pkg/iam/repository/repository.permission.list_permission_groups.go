package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionPermissionGroups(name string) ([]*iam_domain.PermissionGroupFacade, core.Status) {
	r.db.QueryDebug("select * from permission_group_has_permission where permission = ?", name)
	rows, err := r.db.Query(`select permission_group from permission_group_has_permission where permission = ?`, name)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	permissionGroups := make([]*iam_domain.PermissionGroupFacade, 0)
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
