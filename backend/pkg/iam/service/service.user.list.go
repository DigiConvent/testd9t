package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
)

func (service *IAMService) ListUsers(fs *iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], *core.Status) {
	page, status := service.IAMRepository.ListUsers(fs)
	if status.Err() {
		return nil, &status
	}

	return page, &status
}
