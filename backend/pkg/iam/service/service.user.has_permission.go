package iam_service

import (
	"github.com/google/uuid"
)

func (service *IAMService) UserHasPermission(id *uuid.UUID, permission string) bool {
	return service.IAMRepository.UserHasPermission(id, permission)
}
