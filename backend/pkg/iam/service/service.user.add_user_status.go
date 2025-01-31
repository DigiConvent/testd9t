package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/go-playground/validator/v10"
)

func (s *IAMService) AddUserStatus(status *iam_domain.AddUserStatusToUser) *core.Status {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(status)
	if err != nil {
		return &core.Status{Code: 422, Message: err.Error()}
	}

	result := s.IAMRepository.AddUserStatusToUser(status)

	return &result
}
