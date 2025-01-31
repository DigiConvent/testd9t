package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) GetUserByID(id *uuid.UUID) (*iam_domain.UserRead, core.Status) {
	var user = &iam_domain.UserRead{}
	row := r.DB.QueryRow(`SELECT
		id,
		email,
		first_name,
		last_name,
		date_of_birth
	FROM 
		users 
	WHERE 
		id = $1`, id.String())

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
	)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return user, *core.StatusSuccess()
}
