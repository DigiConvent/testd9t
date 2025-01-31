package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) SetEnabled(id *uuid.UUID, enabled bool) core.Status {
	result, err := r.DB.Exec("UPDATE users SET enabled = $1 WHERE id = $2", enabled, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("user not found")
	}
	return *core.StatusNoContent()
}
