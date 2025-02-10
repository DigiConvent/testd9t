package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) SetEnabled(id *uuid.UUID, enabled bool) *core.Status {
	status := s.repository.SetEnabled(id, enabled)
	return &status
}
