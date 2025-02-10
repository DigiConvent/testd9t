package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetUser(id *uuid.UUID) (*iam_domain.UserRead, *core.Status) {
	if id == nil {
		return nil, &core.Status{Code: 422, Message: "ID is required"}
	}

	userRead, status := s.repository.GetUserByID(id)
	if status.Err() {
		return nil, &status
	}

	return userRead, core.StatusSuccess()
}
