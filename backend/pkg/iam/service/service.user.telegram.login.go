// exempt from testing
package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s *IAMService) LoginTelegramUser(body string) (*uuid.UUID, *core.Status) {
	userId, status := s.IAMRepository.VerifyTelegramUser(body)
	if status.Ok() {
		return userId, &status
	}
	return nil, &status
}
