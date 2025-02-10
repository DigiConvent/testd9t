package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetUserProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status) {
	if id == nil {
		return nil, &core.Status{Code: 422, Message: "ID is required"}
	}

	userRead, status := s.repository.GetUserByID(id)
	if status.Err() {
		return nil, &status
	}

	userStatuses, status := s.repository.ListUserStatusesFromUser(id)
	if status.Err() {
		return nil, &status
	}

	userPermissions, status := s.repository.ListUserPermissions(id)
	if status.Err() {
		return nil, &status
	}

	userGroups, status := s.repository.ListUserGroups(id)
	if status.Err() {
		return nil, &status
	}

	return &iam_domain.UserProfile{
		User:        userRead,
		UserStatus:  userStatuses,
		Groups:      userGroups,
		Permissions: userPermissions,
	}, core.StatusSuccess()
}
