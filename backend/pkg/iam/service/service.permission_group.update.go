package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) *core.Status {
	if arg.Parent == uuid.Nil.String() {
		return core.BadRequestError("The super group cannot have descendants")
	}

	currentPermissions, _ := s.repository.ListPermissionGroupPermissions(id)
	impliedPermissions := []string{}
	for _, p := range currentPermissions {
		if p.Implied {
			impliedPermissions = append(impliedPermissions, p.Name)
		}
	}

	revisedPermissions := []string{}
	for _, p := range arg.Permissions {
		if !core_utils.Contains(impliedPermissions, p) {
			revisedPermissions = append(revisedPermissions, p)
		}
	}

	status := s.repository.SetPermissionsForPermissionGroup(id, revisedPermissions)

	if status.Err() {
		return &status
	}

	status = s.repository.UpdatePermissionGroup(id, arg)

	return &status
}
