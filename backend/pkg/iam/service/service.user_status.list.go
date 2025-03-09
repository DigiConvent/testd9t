package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) ListUserStatuses() ([]*iam_domain.UserStatusRead, *core.Status) {
	userStatuses, status := service.repository.ListUserStatuses()
	if status.Err() {
		return nil, &status
	}
	return userStatuses, &status
}
