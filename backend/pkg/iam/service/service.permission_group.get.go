package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status) {
	group, status := s.IAMRepository.GetPermissionGroup(id)

	if status.Err() {
		return nil, &status
	}

	users, status := s.IAMRepository.ListGroupUsers(id)
	if status.Err() {
		return nil, &status
	}

	group.Members = users

	return group, &status
}
