package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) LoginTelegramUser(body, botToken string) (*uuid.UUID, *core.Status) {
	telegramId, status := s.repository.GetTelegramID(body, botToken)
	if status.Err() {
		return nil, &status
	}

	userId, status := s.repository.GetUserByTelegramID(telegramId)
	if status.Err() {
		return nil, &status
	}

	return userId, &status
}
