package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) GetUserByID(id *uuid.UUID) (*iam_domain.UserRead, core.Status) {
	if id == nil {
		return nil, *core.UnprocessableContentError("ID is required")
	}
	var user = &iam_domain.UserRead{}
	row := r.db.QueryRow(`select id, emailaddress, first_name, last_name, enabled from users where id = ?`, id.String())

	err := row.Scan(
		&user.ID,
		&user.Emailaddress,
		&user.FirstName,
		&user.LastName,
		&user.Enabled,
	)
	if err != nil {
		return nil, *core.NotFoundError("User not found with that id")
	}

	return user, *core.StatusSuccess()
}
