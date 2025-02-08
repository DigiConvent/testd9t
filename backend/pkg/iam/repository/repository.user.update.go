package iam_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status {
	if user == nil || id == nil {
		return *core.UnprocessableContentError("User or ID is nil")
	}

	userId := id.String()
	user.Email = strings.ToLower(user.Email)
	result, err := r.DB.Exec(`update users set first_name = ?, last_name = ?, email = ?, date_of_birth = ? where id = ?`, user.FirstName, user.LastName, user.Email, user.DateOfBirth, userId)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("User not found")
	}
	return *core.StatusNoContent()
}
