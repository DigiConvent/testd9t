package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) RemovePermissionFromPermissionGroup(permissionGroupId *uuid.UUID, permission string) *core.Status {
	status := service.repository.RemovePermissionFromPermissionGroup(permissionGroupId, permission)
	return &status
}
