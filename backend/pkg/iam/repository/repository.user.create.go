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
	user.Email = strings.ToLower(user.Email)
	result, err := r.DB.Exec("insert into users (id, first_name, last_name, email, date_of_birth, enabled) values (?, ?, ?, ?, ?, ?)", uid.String(), user.FirstName, user.LastName, user.Email, user.DateOfBirth, false)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, *core.InternalError("Failed to create user")
	}

	return &uid, *core.StatusCreated()
}
