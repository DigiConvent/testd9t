package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissions() ([]*iam_domain.PermissionRead, core.Status) {
	var permissions []*iam_domain.PermissionRead
	rows, err := r.DB.Query(`select name, coalesce(description, name), meta from permissions`)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		permission := &iam_domain.PermissionRead{}
		err = rows.Scan(&permission.Name, &permission.Description, &permission.Meta)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissions = append(permissions, permission)
	}
	return permissions, *core.StatusSuccess()
}
