package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (service *IAMService) ResetPassword(email string) (string, *core.Status) {
	user, status := service.IAMRepository.GetUserByEmail(email)
	if status.Err() {
		return "", &status
	}

	token, status := service.IAMRepository.ResetCredentials(&user.ID)

	return token, &status
}
