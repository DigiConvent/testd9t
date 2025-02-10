package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (service *IAMService) ResetPassword(email string) (string, *core.Status) {
	user, status := service.repository.GetUserByEmail(email)
	if status.Err() {
		return "", &status
	}

	token, status := service.repository.ResetCredentials(&user.ID)

	return token, &status
}
