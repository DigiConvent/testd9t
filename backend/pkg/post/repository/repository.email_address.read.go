package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, core.Status) {
	if id == nil {
		return nil, *core.UnprocessableContentError("ID is required")
	}

	readAddress := &post_domain.EmailAddressRead{}
	err := p.db.QueryRow("select id, name, domain from email_addresses where id = ?", id.String()).Scan(&readAddress.Id, &readAddress.Name, &readAddress.Domain)

	if err != nil {
		return nil, *core.NotFoundError("email address not found")
	}

	return readAddress, *core.StatusSuccess()
}
