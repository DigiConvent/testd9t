package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (iamService *IAMService) DeleteUserStatus(id *uuid.UUID) *core.Status {
	status := iamService.IAMRepository.DeleteUserStatus(id)
	return &status
}
