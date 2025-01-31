package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
	"github.com/jackc/pgtype"
)

func (r *IAMRepository) ListPermissionGroups() ([]*iam_domain.PermissionGroupRead, core.Status) {
	var permissionGroups []*iam_domain.PermissionGroupRead
	rows, err := r.DB.Query(`SELECT
		id,
		name,
		abbr,
		description,
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
		var permissionGroup iam_domain.PermissionGroupRead
		var parent pgtype.UUID
		var id string
		err = rows.Scan(
			&id,
			&permissionGroup.Name,
			&permissionGroup.Abbr,
			&permissionGroup.Description,
			&permissionGroup.IsGroup,
			&permissionGroup.IsNode,
			&parent,
			&permissionGroup.Generated)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		pgId := uuid.MustParse(id)
		permissionGroup.ID = pgId

		if parent.Status != pgtype.Null {
			parent.AssignTo(&permissionGroup.Parent)
		}

		permissionGroups = append(permissionGroups, &permissionGroup)
	}
	return permissionGroups, *core.StatusSuccess()
}
