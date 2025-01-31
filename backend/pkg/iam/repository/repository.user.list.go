package iam_repository

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListUsers() ([]*iam_domain.UserFacade, core.Status) {
	users := []*iam_domain.UserFacade{}
	rows, err := r.DB.Query("SELECT id, name FROM user_facades")
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		user := iam_domain.UserFacade{}

		err := rows.Scan(&user.ID, &user.Name)
		fmt.Println(user.ID, user.Name)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		users = append(users, &user)
	}
	return users, *core.StatusSuccess()
}
