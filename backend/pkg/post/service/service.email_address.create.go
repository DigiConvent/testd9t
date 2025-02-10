package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (s PostService) CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, *core.Status) {
	if credentials == nil {
		return nil, core.UnprocessableContentError("credentials are required")
	}
	id, status := s.repository.CreateEmailAddress(credentials)
	if status.Err() {
		return nil, &status
	}
	return id, &status
}
