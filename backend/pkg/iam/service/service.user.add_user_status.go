package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) AddUserBecameStatus(status *iam_domain.UserBecameStatusWrite) *core.Status {
	result := s.repository.AddUserStatusToUser(status)

	return &result
}
