package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

// UserHasPermission implements IAMServiceInterface.
func (service *IAMService) UserHasPermission(id *uuid.UUID, permission string) (bool, *core.Status) {
	hasPermission, status := service.IAMRepository.UserHasPermission(id, permission)

	return hasPermission, &status
}
