package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, core.Status) {
	id, _ := uuid.NewV7()

	_, err := r.DB.Exec(`INSERT INTO permission_groups (id, name, abbr, description, is_group, is_node, parent) VALUES (?, ?, ?, ?, ?, ?, ?)`, id, arg.Name, arg.Abbr, arg.Description, arg.IsGroup, arg.IsNode, arg.Parent)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return &id, *core.StatusCreated()
}
