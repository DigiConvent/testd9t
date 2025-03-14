package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) CreatePermission(permission *iam_domain.PermissionWrite) *core.Status {
	status := service.repository.CreatePermission(permission)
	return &status
}
