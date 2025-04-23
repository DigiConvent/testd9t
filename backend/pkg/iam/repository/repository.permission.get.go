package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) GetPermission(name string) (*iam_domain.PermissionRead, core.Status) {
	result := r.db.QueryRow("select name, description, meta, generated, archived from permissions where name = ?", name)

	if result.Err() != nil {
		return nil, *core.NotFoundError("iam.permission")
	}

	var permission iam_domain.PermissionRead
	err := result.Scan(&permission.Name, &permission.Description, &permission.Meta, &permission.Generated, &permission.Archived)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return &permission, *core.StatusSuccess()
}
