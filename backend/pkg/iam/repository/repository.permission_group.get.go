package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroupRead, core.Status) {
	if arg == nil {
		return nil, *core.UnprocessableContentError("Permission group ID is required")
	}
	pg := &iam_domain.PermissionGroupRead{}

	row := r.db.QueryRow(`SELECT id, name, abbr, description, is_group, parent, is_node, "generated" FROM permission_groups WHERE id = ?`, arg.String())

	err := row.Scan(
		&pg.ID,
		&pg.Name,
		&pg.Abbr,
		&pg.Description,
		&pg.IsGroup,
		&pg.Parent,
		&pg.IsNode,
		&pg.Generated,
	)

	if err != nil || pg.ID == uuid.Nil {
		return nil, *core.NotFoundError("Permission group not found")
	}

	return pg, *core.StatusSuccess()
}
