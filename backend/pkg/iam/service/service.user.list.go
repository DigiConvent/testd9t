package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) ListUsers() ([]*iam_domain.UserFacade, *core.Status) {
	users, status := s.IAMRepository.ListUsers()
	if status.Err() {
		return nil, &status
	}
	return users, &status
}
