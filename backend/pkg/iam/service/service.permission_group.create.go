package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func (s *IAMService) CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, *core.Status) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(arg)
	if err != nil {
		return nil, core.UnprocessableContentError(err.Error())
	}

	id, status := s.repository.CreatePermissionGroup(arg)
	if status.Err() && status.Code != 201 {
		return nil, &status
	}

	setPermissionsStatus := s.repository.SetPermissionsForPermissionGroup(id, arg.Permissions)

	if setPermissionsStatus.Err() {
		return nil, &setPermissionsStatus
	}

	return id, &status
}
