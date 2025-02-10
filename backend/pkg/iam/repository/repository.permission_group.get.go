package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroupRead, core.Status) {
	if arg == nil {
		return nil, *core.UnprocessableContentError("Permission group ID is required")
	}
	pg := &iam_domain.PermissionGroupRead{}

	row := r.db.QueryRow(`SELECT id, name, abbr, description, is_group, parent, is_node, "generated" FROM permission_groups WHERE id = ?`, arg.String())

	err := row.Scan(
		&pg.ID,
		&pg.Name,
		&pg.Abbr,
		&pg.Description,
		&pg.IsGroup,
		&pg.Parent,
		&pg.IsNode,
		&pg.Generated,
	)

	if err != nil || pg.ID == uuid.Nil {
		return nil, *core.NotFoundError("Permission group not found")
	}

	return pg, *core.StatusSuccess()

	// var permissionGroup = iam_domain.PermissionGroupProfile{
	// 	PermissionGroup: &iam_domain.PermissionGroupRead{},
	// 	Members:         []*iam_domain.UserFacade{},
	// 	Permissions:     []*iam_domain.PermissionFacade{},
	// }

	// rows, err := r.DB.Query(`SELECT
	// 	pg.id,
	// 	pg.name,
	// 	pg.abbr,
	// 	pg.description,
	// 	pg.is_group,
	// 	pg.parent,
	// 	pg.is_node,
	// 	pg."generated",
	// 	p.id,
	// 	p.name
	// FROM permission_groups pg
	// LEFT JOIN permission_group_has_permission pghp ON pg.id = pghp.permission_group
	// LEFT JOIN permissions p ON pghp.permission = p.id
	// LEFT JOIN permission_group_has_user pghu ON pg.id = pghu.permission_group
	// LEFT JOIN user_facades as u ON pghu.user = u.id
	// WHERE pg.id = ?`, arg.String())

	// if err != nil {
	// 	return nil, *core.InternalError(err.Error())
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	permission := &iam_domain.PermissionFacade{}
	// 	err = rows.Scan(
	// 		&permissionGroup.PermissionGroup.ID,
	// 		&permissionGroup.PermissionGroup.Name,
	// 		&permissionGroup.PermissionGroup.Abbr,
	// 		&permissionGroup.PermissionGroup.Description,
	// 		&permissionGroup.PermissionGroup.IsGroup,
	// 		&permissionGroup.PermissionGroup.Parent,
	// 		&permissionGroup.PermissionGroup.IsNode,
	// 		&permissionGroup.PermissionGroup.Generated,
	// 		&permission.ID,
	// 		&permission.Name,
	// 	)
	// 	if err != nil {
	// 		return nil, *core.InternalError(err.Error())
	// 	}

	// 	permissionGroup.Permissions = append(permissionGroup.Permissions, permission)
	// }
	// if err != nil {
	// 	return nil, *core.InternalError(err.Error())
	// }
	// if permissionGroup.PermissionGroup.ID == uuid.Nil {
	// 	return nil, *core.NotFoundError("Permission group not found")
	// }
	// return &permissionGroup, *core.StatusSuccess()
}
