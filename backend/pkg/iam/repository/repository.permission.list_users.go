package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionUsers(name string) ([]*iam_domain.UserFacade, core.Status) {
	rows, err := r.db.Query(`select distinct id, first_name, last_name, status_id, status_name, role_id, role_name from permission_has_users where permission = ?`, name)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	users := make([]*iam_domain.UserFacade, 0)
	for rows.Next() {
		var user iam_domain.UserFacade
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.StatusId, &user.StatusName, &user.RoleId, &user.RoleName)

		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		users = append(users, &user)
	}

	return users, *core.StatusSuccess()
}
