package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreateUserRole(userRole *iam_domain.UserRoleWrite) (*uuid.UUID, core.Status) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	res, err := r.db.Exec("insert into user_roles (id, name, abbr, description) values (?, ?, ?, ?)", id, userRole.Name, userRole.Abbr, userRole.Description)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	d, err := res.RowsAffected()
	if err != nil || d == 0 {
		return nil, *core.InternalError(err.Error())
	}
	r.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{ID: &id, Parent: userRole.Parent})

	return &id, *core.StatusSuccess()
}
