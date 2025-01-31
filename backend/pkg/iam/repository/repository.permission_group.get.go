package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
	"github.com/jackc/pgtype"
)

func (r *IAMRepository) GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroupProfile, core.Status) {
	var permissionGroup = iam_domain.PermissionGroupProfile{
		PermissionGroup: &iam_domain.PermissionGroupRead{},
		Members:         []*iam_domain.UserFacade{},
		Permissions:     []*iam_domain.PermissionFacade{},
	}

	rows, err := r.DB.Query(`SELECT 
		pg.id,
		pg.name,
		pg.abbr,
		pg.description,
		pg.is_group,
		pg.parent,
		pg.is_node,
		pg."generated",
		p.id,
		p.name
	FROM permission_groups pg
	LEFT JOIN permission_group_has_permission pghp ON pg.id = pghp.permission_group
	LEFT JOIN permissions p ON pghp.permission = p.id
	LEFT JOIN permission_group_has_user pghu ON pg.id = pghu.permission_group
	LEFT JOIN user_facades u ON pghu.user = u.id
	WHERE pg.id = $1`, arg.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var pgId pgtype.UUID
		var pgName string
		var pgAbbr string
		var pgDescription string
		var pgIsGroup bool
		var pgParent pgtype.UUID
		var pgIsNode bool
		var pgGenerated bool

		var pId pgtype.UUID
		var pName pgtype.Varchar

		err = rows.Scan(
			&pgId,
			&pgName,
			&pgAbbr,
			&pgDescription,
			&pgIsGroup,
			&pgParent,
			&pgIsNode,
			&pgGenerated,
			&pId,
			&pName,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		if permissionGroup.PermissionGroup.Name == "" {
			permissionGroup.PermissionGroup.Name = pgName
			permissionGroup.PermissionGroup.Abbr = pgAbbr
			permissionGroup.PermissionGroup.Description = pgDescription
			permissionGroup.PermissionGroup.IsGroup = pgIsGroup
			permissionGroup.PermissionGroup.IsNode = pgIsNode
			permissionGroup.PermissionGroup.Generated = pgGenerated
			if pgParent.Status == pgtype.Present {
				var parentId uuid.UUID
				err = pgParent.AssignTo(&parentId)
				if err != nil {
					return nil, *core.InternalError(err.Error())
				}
				permissionGroup.PermissionGroup.Parent = &parentId
			}
		}

		if pId.Status == pgtype.Present {
			permissionGroup.Permissions = append(permissionGroup.Permissions, &iam_domain.PermissionFacade{
				ID:   uuid.Nil,
				Name: pName.String,
			})
		}
	}
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	return &permissionGroup, *core.StatusSuccess()
}
