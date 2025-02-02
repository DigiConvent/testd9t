package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, core.Status) {
	var id string
	row := r.DB.QueryRow(`INSERT INTO permission_groups (name, abbr, description, is_group, is_node, parent) VALUES (?, ?, ?, ?, ?, ?) RETURNING id`, arg.Name, arg.Abbr, arg.Description, arg.IsGroup, arg.IsNode, arg.Parent)
	err := row.Scan(&id)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	uid, _ := uuid.Parse(id)
	return &uid, *core.StatusCreated()
}
