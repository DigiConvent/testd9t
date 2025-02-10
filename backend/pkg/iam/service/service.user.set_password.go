package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) SetUserPassword(id *uuid.UUID, password string) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("ID is required")
	}
	if password == "" {
		return core.UnprocessableContentError("Password is required")
	}

	password, err := hashedPassword(password)
	if err != nil {
		return core.InternalError(err.Error())
	}

	status := service.IAMRepository.SetCredentialPassword(id, password)
	if status.Err() {
		return &status
	}

	return core.StatusSuccess()
}
