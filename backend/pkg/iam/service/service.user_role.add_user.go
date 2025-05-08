package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (s *IAMService) AddUserToUserRole(status *iam_domain.AddRoleToUserWrite) *core.Status {
	if status.Start.IsZero() {
		return core.UnprocessableContentError("Start date is required")
	}
	if status.End.IsZero() {
		return core.UnprocessableContentError("End date is required")
	}
	if status.End.Before(status.Start) {
		status.End, status.Start = status.Start, status.End
	}
	result := s.repository.AddRoleToUser(status)

	return &result
}
