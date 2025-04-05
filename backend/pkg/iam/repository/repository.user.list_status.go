package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserStatusesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasStatusRead, core.Status) {
	var userStatuses []*iam_domain.UserHasStatusRead
	rows, err := r.db.Query(`select
		s.id,
		s.name,
		s.abbr,
		ubs.description,
		ubs.start
	from 
		user_became_status ubs
	join user_status s on ubs.status = s.id
	where 
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
			&userStatus.Start,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userStatuses = append(userStatuses, userStatus)
	}

	return userStatuses, *core.StatusSuccess()
}
