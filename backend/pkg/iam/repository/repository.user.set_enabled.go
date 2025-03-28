package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) SetEnabled(id *uuid.UUID, enabled bool) core.Status {
	if id == nil {
		return *core.UnprocessableContentError("ID is required")
	}
	result, err := r.db.Exec("update users set enabled = ? where id = ?", enabled, id)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("user not found")
	}
	if rowsAffected == 1 && !enabled {
		disableUser(id)
	} else {
		enableUser(id)
	}
	return *core.StatusNoContent()
}
