package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
)

func (service *IAMService) ResetPassword(emailaddress string) (string, *core.Status) {
	user, status := service.repository.GetUserByEmailaddress(emailaddress)
	if status.Err() {
		return "", &status
	}

	token, status := service.repository.ResetCredentials(&user.Id)

	return token, &status
}
