package iam_service

import "github.com/DigiConvent/testd9t/core"

func (s *IAMService) RegisterTelegramUser(telegramId int, email string, code string) *core.Status {
	status := s.IAMRepository.RegisterTelegramUser(telegramId, email, code)

	return &status
}
