package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (service *IAMService) GetUserRoleProfile(id *uuid.UUID) (*iam_domain.UserRoleProfile, *core.Status) {
	userRole, status := service.repository.GetUserRole(id)
	if status.Err() {
		return nil, &status
	}

	users, status := service.repository.ListUserRoleUsers(id, false)
	if status.Err() {
		return nil, &status
	}

	permissionGroup, pgStatus := service.GetPermissionGroupProfile(id)
	if pgStatus.Err() {
		return nil, pgStatus
	}

	return &iam_domain.UserRoleProfile{
		PermissionGroupProfile: permissionGroup,
		UserRole:               userRole,
		History:                users,
	}, &status
}
