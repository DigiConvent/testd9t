package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListStatuses() ([]*iam_domain.UserStatusRead, core.Status) {
	var userStatuses []*iam_domain.UserStatusRead
	rows, err := r.db.Query(`select  id, name, abbr, description, archived from permission_groups where meta = 'status'`)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	for rows.Next() {
		userStatusRead := &iam_domain.UserStatusRead{}
		err = rows.Scan(
			&userStatusRead.Id,
			&userStatusRead.Name,
			&userStatusRead.Abbr,
			&userStatusRead.Description,
			&userStatusRead.Archived,
		)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userStatuses = append(userStatuses, userStatusRead)
	}

	return userStatuses, *core.StatusSuccess()
}
