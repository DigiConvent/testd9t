package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserStatusUsers(arg *uuid.UUID) ([]*iam_domain.UserFacade, core.Status) {
	rows, err := r.DB.Query("SELECT u.id, u.name FROM user_became_status ubs JOIN user_facades u ON ubs.user = u.id WHERE ubs.status = $1 and ubs.active = true", arg.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	users := make([]*iam_domain.UserFacade, 0)

	for rows.Next() {
		var user iam_domain.UserFacade
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		users = append(users, &user)
	}

	return users, *core.StatusSuccess()
}
