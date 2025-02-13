package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, *core.Status) {
	permissionGroups, status := s.repository.ListPermissionGroups()
	if status.Err() {
		return nil, &status
	}
	return permissionGroups, core.StatusSuccess()
}
