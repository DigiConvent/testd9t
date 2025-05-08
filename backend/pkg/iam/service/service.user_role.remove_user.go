package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) RemoveUserFromUserRole(status *iam_domain.AddRoleToUserWrite) *core.Status {
	result := service.repository.RemoveUserFromUserRole(status)

	return &result
}
