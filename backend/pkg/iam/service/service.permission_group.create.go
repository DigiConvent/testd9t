package iam_service

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, *core.Status) {
	if arg == nil {
		return nil, core.UnprocessableContentError("iam.permission_group.create.missing_data")
	}
	if arg.Name == "" {
		return nil, core.UnprocessableContentError("iam.permission_group.create.invalid_name")
	}
	if arg.Parent == nil {
		fmt.Println(arg)
		return nil, core.UnprocessableContentError("iam.permission_group.create.invalid_parent")
	}
	if *arg.Parent == uuid.Nil {
		return nil, core.BadRequestError("The admin role cannot have descendants")
	}

	id, status := s.repository.CreatePermissionGroup(arg)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}

	setPermissionsStatus := s.repository.SetPermissionsForPermissionGroup(id, arg.Permissions)

	if setPermissionsStatus.Err() {
		return nil, &setPermissionsStatus
	}

	return id, &status
}
