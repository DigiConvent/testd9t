package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status) {
	var permissions = make([]*iam_domain.PermissionFacade, 0)
	rows, err := r.db.Query(`
select permission from user_has_permissions where user = ?
`, id.String())
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var permission iam_domain.PermissionFacade
		err := rows.Scan(&permission.Name)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissions = append(permissions, &permission)
	}

	return permissions, *core.StatusSuccess()
}
