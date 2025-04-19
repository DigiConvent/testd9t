package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetUserStatus(id *uuid.UUID) (*iam_domain.UserStatusRead, *core.Status) {
	userStatus, status := s.repository.GetUserStatus(id)
	if status.Code != 200 {
		return nil, &status
	}

	return userStatus, &status
}
