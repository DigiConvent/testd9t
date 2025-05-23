package iam_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, core.Status) {
	uid, err := uuid.NewV7()
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	user.Emailaddress = strings.ToLower(user.Emailaddress)
	result, err := r.db.Exec("insert into users (id, first_name, last_name, emailaddress, enabled) values (?, ?, ?, ?, ?)", uid.String(), user.FirstName, user.LastName, strings.ToLower(user.Emailaddress), false)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, *core.InternalError("Failed to create user")
	}

	return &uid, *core.StatusCreated()
}
