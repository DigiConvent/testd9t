package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListUserRoles() ([]*iam_domain.UserRoleRead, core.Status) {
	var userStatuses []*iam_domain.UserRoleRead
	rows, err := r.db.Query(`select id, name, abbr, description from permission_groups where meta = 'role'`)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	for rows.Next() {
		userRoleRead := &iam_domain.UserRoleRead{}
		err = rows.Scan(
			&userRoleRead.Id,
			&userRoleRead.Name,
			&userRoleRead.Abbr,
			&userRoleRead.Description,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userStatuses = append(userStatuses, userRoleRead)
	}

	return userStatuses, *core.StatusSuccess()
}
