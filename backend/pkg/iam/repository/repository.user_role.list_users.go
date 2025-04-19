package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserRoleUsers(arg *uuid.UUID, now bool) ([]*iam_domain.UserBecameRoleRead, core.Status) {
	rows, err := r.db.Query(`select 
	u.id, 
	u.first_name,
	u.last_name,
	ubr.start,
	ubr.end
	from user_became_role ubr
	join user_facades u on ubr.user = u.id 
	where ubr.role = ?`, arg.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	ubrs := make([]*iam_domain.UserBecameRoleRead, 0)

	for rows.Next() {
		var ubr iam_domain.UserBecameRoleRead = iam_domain.UserBecameRoleRead{
			User: iam_domain.UserFacade{},
		}
		err := rows.Scan(&ubr.User.Id, &ubr.User.FirstName, &ubr.User.LastName, &ubr.Start, &ubr.End)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		ubrs = append(ubrs, &ubr)
	}

	return ubrs, *core.StatusSuccess()
}
