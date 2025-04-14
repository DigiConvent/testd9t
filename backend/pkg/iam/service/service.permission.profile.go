package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) GetPermissionProfile(name string) (*iam_domain.PermissionProfile, *core.Status) {
	permission, status := service.repository.GetPermission(name)
	if status.Err() {
		return nil, &status
	}

	permisionGroups, status := service.repository.ListPermissionPermissionGroups(name)
	if status.Err() {
		return nil, &status
	}

	permissionUsers, status := service.repository.ListPermissionUsers(name)
	if status.Err() {
		return nil, &status
	}

	descendantPermissions, status := service.repository.ListPermissionDescendants(name)
	if status.Err() {
		return nil, &status
	}

	return &iam_domain.PermissionProfile{
		Permission:       permission,
		PermissionGroups: permisionGroups,
		Users:            permissionUsers,
		Descendants:      descendantPermissions,
	}, core.StatusSuccess()
}
