package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) *core.Status {
	status := service.repository.AddUserToPermissionGroup(permissionGroup, userId)
	return &status
}
