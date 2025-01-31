package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) AddUserStatusToUser(d *iam_domain.AddUserStatusToUser) core.Status {
	result, err := r.DB.Exec(`INSERT INTO user_became_status ("user", status, date, description) VALUES ($1, $2, $3, $4)`, d.UserID, d.StatusID, d.When, d.Description)
	if err != nil {
		return *core.InternalError(err.Error())
	}
	lastInsertId, _ := result.LastInsertId()

	if lastInsertId == 0 {
		return *core.InternalError("Failed to add user status")
	}
	return *core.StatusSuccess()
}
