package iam_repository

import (
	"encoding/json"
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, core.Status) {
	id, _ := uuid.NewV7()

	var parent *string
	if arg.Parent != "" {
		parent = &arg.Parent
	}

	_, err := r.db.Exec(`insert into permission_groups (id, name, abbr, description, is_group, is_node, parent) values (?, ?, ?, ?, ?, ?, ?)`, id, arg.Name, arg.Abbr, arg.Description, arg.IsGroup, arg.IsNode, parent)

	if err != nil {
		fmt.Println(arg)
		feedback, err2 := json.Marshal(&arg)
		if err2 != nil {
			log.Error(err2.Error())
			return nil, *core.InternalError(err2.Error())
		}
		return nil, *core.InternalError(err.Error() + string(feedback))
	}

	return &id, *core.StatusCreated()
}
