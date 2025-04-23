package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) SetUserPassword(id *uuid.UUID, rawPassword string) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("iam.user.set_password.missing_user")
	}
	if rawPassword == "" {
		return core.UnprocessableContentError("iam.user.set_password.empty")
	}

	password, err := hashedPassword(rawPassword)
	if err != nil {
		return core.InternalError(err.Error())
	}

	status := service.repository.SetCredentialPassword(id, password)
	if status.Err() {
		return &status
	}

	return core.StatusNoContent()
}
