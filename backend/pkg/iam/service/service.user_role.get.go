package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetUserRole(id *uuid.UUID) (*iam_domain.UserRoleRead, *core.Status) {
	UserRole, status := s.repository.GetUserRole(id)
	if status.Code != 200 {
		return nil, &status
	}

	return UserRole, &status
}
