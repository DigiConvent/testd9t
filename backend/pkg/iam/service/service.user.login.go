package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) LoginUser(email string, password string) (*uuid.UUID, *core.Status) {
	password, err := hashedPassword(password)
	if err != nil {
		return nil, core.InternalError(err.Error())
	}
	userId, status := service.repository.GetCredentials(email, password)

	if status.Err() {
		return nil, &status
	}

	return userId, core.StatusSuccess()
}
