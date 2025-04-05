package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) AddUserRole(status *iam_domain.AddUserRoleToUser) *core.Status {
	result := s.repository.AddUserRoleToUser(status)

	return &result
}
