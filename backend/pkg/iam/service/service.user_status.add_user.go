package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) AddUserToUserStatus(status *iam_domain.UserBecameStatusWrite) *core.Status {
	result := s.repository.AddStatusToUser(status)

	return &result
}
