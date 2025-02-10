package post_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, core.Status) {
	if credentials == nil {
		return nil, *core.UnprocessableContentError("credentials are required")
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	result, err := p.db.Exec("insert into email_addresses (id, name, domain) values (?, ?, ?)", id.String(), strings.ToLower(credentials.Name), strings.ToLower(credentials.Domain))

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, *core.InternalError("Failed to create email address")
	}

	return &id, *core.StatusCreated()
}
