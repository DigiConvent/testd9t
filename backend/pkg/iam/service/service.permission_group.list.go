package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) ListPermissionGroups() ([]*iam_domain.PermissionGroupRead, *core.Status) {
	permissionGroups, status := s.IAMRepository.ListPermissionGroups()
	if status.Err() {
		return nil, &status
	}
	return permissionGroups, nil
}
