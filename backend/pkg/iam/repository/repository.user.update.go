package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status {
	if user == nil || id == nil {
		return *core.UnprocessableContentError("User or ID is nil")
	}

	userId := id.String()
	result, err := r.db.Exec(`update users set first_name = coalesce(nullif(?, ''), first_name), last_name = coalesce(nullif(?, ''), last_name), emailaddress = coalesce(nullif(?, ''), emailaddress), date_of_birth = coalesce(nullif(?, ''), date_of_birth) where id = ?`, user.FirstName, user.LastName, user.Emailaddress, user.DateOfBirth, userId)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("User not found")
	}
	return *core.StatusNoContent()
}
