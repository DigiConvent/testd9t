package iam_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) CreatePermission(permission *iam_domain.PermissionWrite) core.Status {
	row := r.db.QueryRow("select 1 from permissions where name = ?", permission.Name)

	var exists int
	row.Scan(&exists) // err will be non-null if the query fails
	if exists == 1 {
		return *core.ConflictError("Permission already exists")
	}

	if strings.Contains(permission.Name, ".") {
		segments := strings.Split(permission.Name, ".")
		r.CreatePermission(&iam_domain.PermissionWrite{Name: strings.Join(segments[:len(segments)-1], ".")})
	}

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
