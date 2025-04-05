package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdateUserRole(id *uuid.UUID, arg *iam_domain.UserRoleWrite) *core.Status {
	status := s.repository.UpdateUserRole(id, arg)
	return &status
}
