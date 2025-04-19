package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroupRead, *core.Status) {
	read, status := s.repository.GetPermissionGroup(id)

	if status.Err() {
		return nil, &status
	}

	read.Permissions, status = s.repository.ListPermissionGroupPermissions(&read.Id)

	return read, &status
}
