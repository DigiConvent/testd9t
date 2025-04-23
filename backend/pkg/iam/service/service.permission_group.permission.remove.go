package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) RemovePermissionFromPermissionGroup(permissionGroupId *uuid.UUID, permission string) *core.Status {
	if permissionGroupId == nil {
		return core.NotFoundError("iam.permission_group.missing_permission_group")
	}
	if permission == "" {
		return core.NotFoundError("iam.permission_group.missing_permission")
	}
	p, status := service.repository.GetPermission(permission)
	if p == nil && status.Code == 404 {
		return core.NotFoundError("iam.permission_group.missing_permission")
	}

	status = service.repository.RemovePermissionFromPermissionGroup(permissionGroupId, permission)
	return &status
}
