package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserStatusesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasStatusRead, core.Status) {
	var userStatuses []*iam_domain.UserHasStatusRead
	rows, err := r.DB.Query(`SELECT
		s.id,
		s.name,
		s.abbr,
		ubs.description,
		ubs.date
	FROM 
		user_became_status ubs
	JOIN user_status s ON ubs.status = s.id
	WHERE 
		ubs."user" = ?`, id.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {

		userStatus := &iam_domain.UserHasStatusRead{}

		err := rows.Scan(
			&userStatus.ID,
			&userStatus.Name,
			&userStatus.Abbr,
			&userStatus.Description,
			&userStatus.Date,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userStatuses = append(userStatuses, userStatus)
	}

	return userStatuses, *core.StatusSuccess()
}
