package iam_repository

import (
	"encoding/json"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

var allowedSortBy = []string{"last_name", "first_name", "emailaddress"}

func (r *IAMRepository) ListUsers(filter *iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], core.Status) {
	sortClause := ""
	if filter != nil && core_utils.Contains(allowedSortBy, filter.Sort.Field) {
		sortClause = "order by " + filter.Sort.Field
		if !filter.Sort.Asc {
			sortClause += " desc"
		} else {
			sortClause += " asc"
		}
	}

	filterClause := ""

	users := []*iam_domain.UserFacade{}
	rows, err := r.db.Query("select id, first_name, last_name, status_id, status_name, roles from user_facades " + sortClause)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		user := iam_domain.UserFacade{}

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

	pageNumber := 1
	if filter != nil {
		pageNumber = filter.Page
	}

	itemsPerPage := 10
	if filter != nil {
		itemsPerPage = filter.ItemsPerPage
	}
	var page = &pagination.Page[*iam_domain.UserFacade]{
		Items:        users,
		Page:         pageNumber,
		ItemsPerPage: itemsPerPage,
	}

	err = r.db.QueryRow("select count(*) from user_facades " + filterClause).Scan(&page.ItemsCount)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return page, *core.StatusSuccess()
}
