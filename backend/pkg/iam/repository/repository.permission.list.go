package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (r *IAMRepository) ListPermissions() ([]*iam_domain.PermissionRead, core.Status) {
	var permissions []*iam_domain.PermissionRead
	rows, err := r.DB.Query(`SELECT id, name, coalesce(description, name), meta FROM permissions`)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var permission iam_domain.PermissionRead
		var id string
		err = rows.Scan(&id, &permission.Name, &permission.Description, &permission.Meta)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		userId := uuid.MustParse(id)
		permission.ID = userId
		permissions = append(permissions, &permission)
	}
	return permissions, *core.StatusSuccess()
}
