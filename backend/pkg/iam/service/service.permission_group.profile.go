package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (service *IAMService) GetPermissionGroupProfile(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status) {
	profile := &iam_domain.PermissionGroupProfile{}
	group, status := service.repository.GetPermissionGroup(id)
	if status.Err() {
		return nil, &status
	}

	profile.PermissionGroup = group
	users, status := service.repository.ListGroupUsers(id)
	if status.Err() {
		log.Error(status.Message)
		return nil, &status
	}
	profile.Members = users

	permissionGroups, status := service.repository.ListPermissionGroupAncestors(id)
	if status.Err() {
		return nil, &status
	}
	profile.Ancestors = permissionGroups

	permissionGroups, status = service.repository.ListPermissionGroupDescendants(id)
	if status.Err() {
		return nil, &status
	}
	profile.Descendants = permissionGroups

	permissions, status := service.repository.ListPermissionGroupPermissions(id)
	if status.Err() {
		return nil, &status
	}
	profile.Permissions = permissions

	return profile, &status
}
