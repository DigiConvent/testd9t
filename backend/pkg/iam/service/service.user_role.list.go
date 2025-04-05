package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) ListUserRoles() ([]*iam_domain.UserRoleRead, *core.Status) {
	UserRolees, status := service.repository.ListUserRoles()
	if status.Err() {
		return nil, &status
	}
	return UserRolees, &status
}
