package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserRolesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasRoleRead, core.Status) {
	var userRoles []*iam_domain.UserHasRoleRead
	rows, err := r.db.Query(`select
		s.id,
		s.name,
		s.abbr,
		ubr.description,
		ubr.start
	from 
		user_became_role ubr
	join user_roles s on ubr.status = s.id
	where 
		ubr."user" = ?`, id.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		userRole := &iam_domain.UserHasRoleRead{}

		err := rows.Scan(
			&userRole.Id,
			&userRole.Name,
			&userRole.Abbr,
			&userRole.Comment,
			&userRole.Start,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userRoles = append(userRoles, userRole)
	}

	return userRoles, *core.StatusSuccess()
}
