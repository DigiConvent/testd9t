package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) UpdateEmailAddress(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) core.Status {
	panic("unimplemented")
}
