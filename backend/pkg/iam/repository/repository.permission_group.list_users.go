package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, core.Status) {
	if groupId == nil {
		return nil, *core.UnprocessableContentError("Group ID is required")
	}

	var users = make([]*iam_domain.UserFacade, 0)

	rows, err := r.db.Query(`select "user" as id, first_name, last_name, implied from permission_group_has_users where root = ?`, groupId.String())

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user iam_domain.UserFacade
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Implied)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		users = append(users, &user)
	}

	return users, *core.StatusSuccess()
}
