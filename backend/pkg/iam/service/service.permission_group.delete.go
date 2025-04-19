package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) DeletePermissionGroup(id *uuid.UUID, generated bool) *core.Status {
	status := s.repository.DeletePermissionGroup(id, generated)
	return &status
}
