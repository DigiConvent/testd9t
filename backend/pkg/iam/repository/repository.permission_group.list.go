package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, core.Status) {
	var permissionGroups []*iam_domain.PermissionGroupFacade
	rows, err := r.db.Query(`select id, name, abbr, is_group, is_node, meta, parent, "generated" from permission_groups`)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		permissionGroup := iam_domain.PermissionGroupFacade{}

		var parentId *string

		err = rows.Scan(
			&permissionGroup.ID,
			&permissionGroup.Name,
			&permissionGroup.Abbr,
			&permissionGroup.IsGroup,
			&permissionGroup.IsNode,
			&permissionGroup.Meta,
			&parentId, // there is an entry in the database (admin) where the id is 00000000-0000-0000-0000-000000000000 which equals to uuid.Nil
			// if there is no parent id, the parent id will be nil and thus implies that the group inherits all the permissions from admin
			&permissionGroup.Generated,
		)

		if parentId != nil {
			parsedParentId, err := uuid.Parse(*parentId)
			if err == nil {
				permissionGroup.Parent = &parsedParentId
			}
		}

		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissionGroups = append(permissionGroups, &permissionGroup)
	}
	return permissionGroups, *core.StatusSuccess()
}
