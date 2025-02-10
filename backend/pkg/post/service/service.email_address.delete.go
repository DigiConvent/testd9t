package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (s PostService) DeleteEmailAddress(id *uuid.UUID) *core.Status {
	if id == nil {
		return core.UnprocessableContentError("ID is required")
	}
	status := s.repository.DeleteEmailAddress(id)

	return &status
}
