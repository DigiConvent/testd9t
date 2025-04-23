package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/google/uuid"
)

func (s *IAMService) CreateUserRole(arg *iam_domain.UserRoleWrite) (*uuid.UUID, *core.Status) {
	if arg == nil {
		return nil, core.UnprocessableContentError("iam.user_role.create.no_data")
	}
	if arg.Name == "" {
		return nil, core.UnprocessableContentError("iam.user_role.create.invalid_name")
	}
	if arg.Parent == nil {
		return nil, core.UnprocessableContentError("iam.user_role.create.missing_parent")
	}

	id, status := s.repository.CreateUserRole(arg)
	if status.Err() {
		return nil, &status
	}
	return id, &status
}
