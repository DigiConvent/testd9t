package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) GetTelegramRegistrationCode(userId *uuid.UUID) (string, *core.Status) {
	code, status := s.repository.GetTelegramRegistrationCode(userId)

	return code, &status
}
