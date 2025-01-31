package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetUserStatus(arg *uuid.UUID) (*iam_domain.UserStatusRead, core.Status) {
	row := r.DB.QueryRow("SELECT id, name, abbr, description, archived FROM user_status WHERE id = $1", arg)

	var userStatus iam_domain.UserStatusRead
	err := row.Scan(&userStatus.ID, &userStatus.Name, &userStatus.Abbr, &userStatus.Description, &userStatus.Archived)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return &userStatus, *core.StatusSuccess()
}
