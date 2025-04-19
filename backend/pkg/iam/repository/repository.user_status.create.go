package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreateUserStatus(userStatus *iam_domain.UserStatusWrite) (*uuid.UUID, core.Status) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	res, err := r.db.Exec("insert into user_status (id, name, abbr, description, archived) values (?, ?, ?, ?, ?)", id, userStatus.Name, userStatus.Abbr, userStatus.Description, userStatus.Archived)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	d, err := res.RowsAffected()
	if err != nil || d == 0 {
		return nil, *core.InternalError(err.Error())
	}
	status := r.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{Id: &id, Parent: userStatus.Parent})
	if status.Err() {
		return nil, status
	}

	return &id, *core.StatusSuccess()
}
