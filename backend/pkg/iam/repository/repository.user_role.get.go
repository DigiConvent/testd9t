package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetUserRole(arg *uuid.UUID) (*iam_domain.UserRoleRead, core.Status) {
	row := r.db.QueryRow("select id, name, abbr, description from user_roles where id = ?", arg)

	var userRole iam_domain.UserRoleRead
	err := row.Scan(&userRole.ID, &userRole.Name, &userRole.Abbr, &userRole.Description)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return &userRole, *core.StatusSuccess()
}
