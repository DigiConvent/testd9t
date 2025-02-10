package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, core.Status) {
	var permissionGroups []*iam_domain.PermissionGroupFacade
	rows, err := r.db.Query(`SELECT
		id,
		name,
		abbr,
		is_group,
		is_node,
		parent,
		"generated"
	FROM permission_groups`)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		permissionGroup := iam_domain.PermissionGroupFacade{}

		err = rows.Scan(
			&permissionGroup.ID,
			&permissionGroup.Name,
			&permissionGroup.Abbr,
			&permissionGroup.IsGroup,
			&permissionGroup.IsNode,
			&permissionGroup.Parent,
			&permissionGroup.Generated,
		)

		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissionGroups = append(permissionGroups, &permissionGroup)
	}
	return permissionGroups, *core.StatusSuccess()
}
