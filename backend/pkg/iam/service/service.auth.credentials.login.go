package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (service *IAMService) LoginUser(emailaddress string, rawPassword string) (*uuid.UUID, *core.Status) {
	userId, hashedPassword, status := service.repository.GetCredentials(emailaddress)

	if status.Err() {
		return nil, &status
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword)) != nil {
		return nil, core.UnauthorizedError("Invalid credentials")
	}

	return userId, core.StatusSuccess()
}
