package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (s PostService) UpdateEmailAddresses(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) *core.Status {
	if credentials == nil || id == nil {
		return core.UnprocessableContentError("credentials are required")
	}

	status := s.repository.UpdateEmailAddress(id, credentials)

	return &status
}
