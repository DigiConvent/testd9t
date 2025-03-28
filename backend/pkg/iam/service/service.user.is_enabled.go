package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (service *IAMService) IsEnabled(id *uuid.UUID) (bool, *core.Status) {
	enabled, status := service.repository.IsEnabled(id)
	return enabled, &status
}
