package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (service *IAMService) DeletePermission(name string) *core.Status {
	status := service.repository.DeletePermission(name)
	return &status
}
