package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status) {
	var permissions = make([]*iam_domain.PermissionFacade, 0)
	rows, err := r.DB.Query(`
with recursive relevant_groups as (
	select permission_group from permission_group_has_user where "user" = ?
	union all
	select child.id as permission_group from permission_groups child
		inner join relevant_groups s on s.permission_group = child.parent
)
select distinct p.name from relevant_groups rg 
join permission_group_has_permission pghp on pghp.permission_group = rg.permission_group
join permissions p on pghp.permission = p.id
`, id.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var permission iam_domain.PermissionFacade
		err := rows.Scan(&permission.Name)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissions = append(permissions, &permission)
	}

	return permissions, *core.StatusSuccess()
}
