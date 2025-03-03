package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) SetUserPassword(id *uuid.UUID, rawPassword string) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("ID is required")
	}
	if rawPassword == "" {
		return core.UnprocessableContentError("Password is required")
	}

	password, err := hashedPassword(rawPassword)
	if err != nil {
		return core.InternalError(err.Error())
	}

	status := service.repository.SetCredentialPassword(id, password)
	if status.Err() {
		return &status
	}

	return core.StatusSuccess()
}
