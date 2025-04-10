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

	row := r.db.QueryRow(`select id, name, abbr, description, is_group, parent, is_node, meta, "generated" from permission_groups where id = ?`, arg.String())

	var meta *string
	err := row.Scan(&pg.ID, &pg.Name, &pg.Abbr, &pg.Description, &pg.IsGroup, &pg.Parent, &pg.IsNode, &meta, &pg.Generated)

	if pg.Parent == nil || *pg.Parent == uuid.Nil {
		pg.Parent = nil
	}

	if meta == nil {
		pg.Meta = ""
	} else {
		pg.Meta = *meta
	}

	if err != nil {
		return nil, *core.NotFoundError("permission group " + arg.String() + " not found: " + err.Error())
	}

	return pg, *core.StatusSuccess()
}
