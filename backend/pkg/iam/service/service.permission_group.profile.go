package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (service *IAMService) GetPermissionGroupProfile(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status) {

	profile := &iam_domain.PermissionGroupProfile{}

	group, status := service.IAMRepository.GetPermissionGroup(id)

	if status.Err() {
		return nil, &status
	}

	profile.PermissionGroup = group

	users, status := service.IAMRepository.ListGroupUsers(id)
	if status.Err() {
		return nil, &status
	}

	profile.Members = users

	permissionGroups, status := service.IAMRepository.ListPermissionGroupPermissionGroups(id)

	if status.Err() {
		return nil, &status
	}

	profile.PermissionGroups = permissionGroups

	permissions, status := service.IAMRepository.ListPermissionGroupPermissions(id)

	if status.Err() {
		return nil, &status
	}

	profile.Permissions = permissions

	return profile, &status
}
