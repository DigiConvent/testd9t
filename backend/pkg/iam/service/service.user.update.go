package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) *core.Status {
	status := s.repository.UpdateUser(id, user)
	return &status
}
