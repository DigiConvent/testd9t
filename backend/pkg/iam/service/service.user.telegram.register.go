package iam_service

import "github.com/DigiConvent/testd9t/core"

func (s *IAMService) RegisterTelegramUser(telegramId int, emailaddress string, code string) *core.Status {
	status := s.repository.RegisterTelegramUser(telegramId, emailaddress, code)

	return &status
}
