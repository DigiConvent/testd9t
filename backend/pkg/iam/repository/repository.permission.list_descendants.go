package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (r *IAMRepository) ListPermissionDescendants(name string) ([]*iam_domain.PermissionFacade, core.Status) {
	rows, err := r.db.Query(`select name, description, meta from permissions where name like ?`, name+".%")

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	defer rows.Close()

	permissions := make([]*iam_domain.PermissionFacade, 0)
	for rows.Next() {
		permission := &iam_domain.PermissionFacade{Implied: false}
		err := rows.Scan(&permission.Name, &permission.Description, &permission.Meta)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		permissions = append(permissions, permission)
	}
	return permissions, *core.StatusSuccess()
}
