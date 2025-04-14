package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) AddPermissionToPermissionGroup(permissionGroupId *uuid.UUID, permission string) *core.Status {
	status := service.repository.AddPermissionToPermissionGroup(permissionGroupId, permission)
	return &status
}
