package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (s PostService) ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, *core.Status) {
	if id == nil {
		return nil, core.UnprocessableContentError("PostService requires an ID")
	}
	address, status := s.repository.ReadEmailAddress(id)
	return address, &status
}
