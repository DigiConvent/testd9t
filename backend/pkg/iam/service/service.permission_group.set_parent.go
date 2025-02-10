package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status {
	status := s.repository.SetParentPermissionGroup(arg)
	return &status
}
