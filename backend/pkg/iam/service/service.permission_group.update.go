package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) *core.Status {
	if arg.Parent == uuid.Nil.String() {
		return core.BadRequestError("The admin role cannot have descendants")
	}

	status := s.repository.SetPermissionsForPermissionGroup(id, arg.Permissions)

	if status.Err() {
		return &status
	}

	status = s.repository.UpdatePermissionGroup(id, arg)

	return &status
}
