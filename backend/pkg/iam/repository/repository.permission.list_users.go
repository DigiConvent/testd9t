package iam_repository

import (
	"encoding/json"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionUsers(name string) ([]*iam_domain.UserFacade, core.Status) {
	rows, err := r.db.Query(`select distinct id, first_name, last_name, status_id, status_name, roles from permission_has_users where permission = ?`, name)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	users := make([]*iam_domain.UserFacade, 0)
	for rows.Next() {
		var user iam_domain.UserFacade
		roles := ""
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.StatusID, &user.StatusName, &roles)

		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		if roles != "" {
			json.Unmarshal([]byte(roles), &user.Roles)
		}

		users = append(users, &user)
	}

	return users, *core.StatusSuccess()
}
