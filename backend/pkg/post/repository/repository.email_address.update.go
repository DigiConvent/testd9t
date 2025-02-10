package post_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) UpdateEmailAddress(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) core.Status {
	if credentials == nil || id == nil {
		return *core.UnprocessableContentError("credentials are required")
	}

	result, err := p.db.Exec("update email_addresses set name = ?, domain = ? where id = ?", strings.ToLower(credentials.Name), strings.ToLower(credentials.Domain), id.String())

	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.NotFoundError("email address not found")
	}

	return *core.StatusNoContent()
}
