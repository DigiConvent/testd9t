package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) UpdateUserStatus(id *uuid.UUID, arg *iam_domain.UserStatusWrite) core.Status {
	result, err := r.DB.Exec("UPDATE user_status SET name = ?, abbr = ?, description = ?, archived = ? WHERE id = ?", arg.Name, arg.Abbr, arg.Description, arg.Archived, id.String())
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("User status not found")
	}
	return *core.StatusNoContent()
}
