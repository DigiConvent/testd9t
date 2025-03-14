package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) CreatePermission(permission *iam_domain.PermissionWrite) core.Status {
	result, err := r.db.Exec("insert into permissions (name, description, meta) values (?, ?, ?)", permission.Name, permission.Description, permission.Meta)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.InternalError("Failed to create permission")
	}

	return *core.StatusCreated()
}
