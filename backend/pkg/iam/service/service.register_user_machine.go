package iam_service

import "github.com/DigiConvent/testd9t/core"

func (s *IAMService) RegisterUserMachine(userId string, macAddress string) *core.Status {
	// this is called when a user signs in with a certificate but the certificate hasn't been registered
	// in the database yet. This is a security measure to prevent unauthorized access from multiple devices
	// a certificate can only be registered to one device

	return &core.Status{}
}
