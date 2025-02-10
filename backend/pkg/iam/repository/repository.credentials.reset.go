package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ResetCredentials(id *uuid.UUID) (string, core.Status) {
	if id == nil {
		return "", *core.UnprocessableContentError("ID is required")
	}
	r.db.Exec("delete from reset_credentials_requests where user = ?", id.String())
	// this can fail, I don't care

	randomString, err := uuid.NewV7()
	if err != nil {
		return "", *core.InternalError(err.Error())
	}
	result, err := r.db.Exec("insert into reset_credentials_requests (user, token) values (?, ?)", id.String(), randomString.String())

	if err != nil {
		return "", *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return "", *core.InternalError("Failed to create reset credentials request")
	}

	return randomString.String(), *core.StatusSuccess()
}
