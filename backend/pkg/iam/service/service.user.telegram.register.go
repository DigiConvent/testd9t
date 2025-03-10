// exempt from testing
package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) ConnectTelegramUser(initData, botToken string, userId *uuid.UUID) *core.Status {
	telegramId, status := s.repository.GetTelegramID(initData, botToken)
	if status.Err() {
		return &status
	}

	status = s.repository.RegisterTelegramUser(*telegramId, userId)
	if status.Err() {
		return &status
	}

	return &status
}
