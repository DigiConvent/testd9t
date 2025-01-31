package iam_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status {
	userId := id.String()
	user.Email = strings.ToLower(user.Email)
	result, err := r.DB.Exec(`UPDATE users SET first_name = $2, last_name = $3, email = $4, date_of_birth = $5 WHERE id = $1`, userId, user.FirstName, user.LastName, user.Email, user.DateOfBirth)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("User not found")
	}
	return *core.StatusNoContent()
}
