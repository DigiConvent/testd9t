package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) CreateUserStatus(arg *iam_domain.UserStatusWrite) (*uuid.UUID, *core.Status) {
	id, status := s.repository.CreateStatus(arg)
	if status.Err() {
		return nil, &status
	}
	return id, &status
}
