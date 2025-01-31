package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) DeletePermissionGroup(id *uuid.UUID) *core.Status {
	status := s.IAMRepository.DeletePermissionGroup(id)
	return &status
}
