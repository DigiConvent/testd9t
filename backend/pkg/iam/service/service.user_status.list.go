package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) ListUserStatuses(fs *iam_domain.UserFilterSort) ([]*iam_domain.UserStatusRead, *core.Status) {
	userStatuses, status := service.IAMRepository.ListUserStatuses()
	if status.Err() {
		return nil, &status
	}
	return userStatuses, &status
}
