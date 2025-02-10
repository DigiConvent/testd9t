package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) ListPermissions() ([]*iam_domain.PermissionRead, *core.Status) {
	permissions, status := service.repository.ListPermissions()
	if status.Err() {
		return nil, &status
	}
	return permissions, &status
}
