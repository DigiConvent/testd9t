package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, core.Status) {
	var users = make([]*iam_domain.UserFacade, 0)
	rows, err := r.DB.Query(`SELECT "user" as id, name, implied from permission_group_has_users where root = $1`, groupId.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user iam_domain.UserFacade
		err := rows.Scan(&user.ID, &user.Name, &user.Implied)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		users = append(users, &user)
	}

	return users, *core.StatusSuccess()
}
