package iam_repository

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserStatusUsers(arg *uuid.UUID) ([]*iam_domain.UserBecameStatusRead, core.Status) {

	rows, err := r.db.Query(`select 
		ubs.status,
		uf.id, 
		uf.first_name,
		uf.last_name,
		ubs.start,
		ubs.end,
		ubs.description
	from user_became_status ubs
		left join user_facades uf on ubs.user = uf.id
	where ubs.status = ? and ubs.start <= datetime('now', 'localtime') and (ubs."end" >= datetime('now', 'localtime') or ubs."end" is null)`, arg.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	usersWhoBecameStatus := make([]*iam_domain.UserBecameStatusRead, 0)

	for rows.Next() {
		var ubs iam_domain.UserBecameStatusRead = iam_domain.UserBecameStatusRead{
			User: iam_domain.UserFacade{},
		}
		var startTime, endTime *string
		err := rows.Scan(
			&ubs.UserStatus,
			&ubs.User.Id,
			&ubs.User.FirstName,
			&ubs.User.LastName,
			&startTime,
			&endTime,
			&ubs.Description,
		)

		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		if endTime != nil {
			end, err := time.Parse(time.RFC3339, *endTime)
			if err != nil {
				return nil, *core.InternalError(err.Error())
			}

			ubs.End = &end
		}

		if *startTime != "" {
			start, err := time.Parse(time.RFC3339, *startTime)
			if err != nil {
				return nil, *core.InternalError(err.Error())
			}
			ubs.Start = start
		}

		usersWhoBecameStatus = append(usersWhoBecameStatus, &ubs)
	}

	return usersWhoBecameStatus, *core.StatusSuccess()
}
