package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdateUserStatus(id *uuid.UUID, arg *iam_domain.UserStatusWrite) *core.Status {
	status := s.repository.UpdateUserStatus(id, arg)
	return &status
}
