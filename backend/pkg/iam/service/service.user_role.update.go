package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) UpdateUserRole(id *uuid.UUID, arg *iam_domain.UserRoleWrite) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("iam.user_role.update.missing_user_role")
	}
	if arg == nil {
		return core.UnprocessableContentError("iam.user_role.update.missing_data")
	}
	if arg.Name == "" {
		return core.UnprocessableContentError("iam.user_role.update.invalid_name")
	}
	if arg.Parent == nil {
		return core.UnprocessableContentError("iam.user_role.update.missing_parent")
	}
	status := s.repository.UpdateUserRole(id, arg)
	return &status
}
