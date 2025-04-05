package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserRoleUsers(arg *uuid.UUID) ([]*iam_domain.UserFacade, core.Status) {
	rows, err := r.db.Query(`select 
	u.id, 
	u.name 
	from user_became_status ubs
	join user_facades u on ubs.user = u.id 
	where ubs.status = ?`, arg.String())
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
