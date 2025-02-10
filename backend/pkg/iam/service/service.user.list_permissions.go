package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, *core.Status) {
	permissions, status := s.repository.ListUserPermissions(id)
	if status.Err() {
		return nil, &status
	}
	return permissions, core.StatusSuccess()
}
