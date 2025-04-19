package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (service *IAMService) GetUserStatusProfile(id *uuid.UUID) (*iam_domain.UserStatusProfile, *core.Status) {
	var profile iam_domain.UserStatusProfile

	permissionGroupProfile, permissionGroupStatus := service.GetPermissionGroupProfile(id)
	if permissionGroupStatus.Err() {
		return nil, permissionGroupStatus
	}
	profile.PermissionGroupProfile = permissionGroupProfile

	userStatus, status := service.repository.GetUserStatus(id)
	if status.Code != 200 {
		return nil, &status
	}
	profile.UserStatus = userStatus

	users, status := service.repository.ListUserStatusUsers(id)
	if status.Code != 200 {
		return nil, &status
	}
	profile.History = users

	return &profile, core.StatusSuccess()
}
