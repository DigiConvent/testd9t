package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

var allowedSortBy = []string{"last_name", "first_name", "email", "date_of_birth"}

func (r *IAMRepository) ListUsers(filter *iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], core.Status) {
	sortClause := ""
	if filter != nil && core_utils.Contains[string](allowedSortBy, filter.Sort.Field) {
		sortClause = "order by " + filter.Sort.Field
		if !filter.Sort.Asc {
			sortClause += " desc"
		} else {
			sortClause += " asc"
		}
	}

	users := []*iam_domain.UserFacade{}
	rows, err := r.db.Query("select id, name from user_facades " + sortClause)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		user := iam_domain.UserFacade{}

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		users = append(users, &user)
	}

	return &pagination.Page[*iam_domain.UserFacade]{
		Items: users,
	}, *core.StatusSuccess()
}

// 	if err != nil {
// 		return nil, *core.InternalError(err.Error())
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		user := iam_domain.UserFacade{}

// 		err := rows.Scan(&user.ID, &user.Name)
// 		fmt.Println(user.ID, user.Name)
// 		if err != nil {
// 			return nil, *core.InternalError(err.Error())
// 		}

// 		users = append(users, &user)
// 	}
// 	return users, *core.StatusSuccess()
// }
