package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/google/uuid"
)

func (p PostRepository) DeleteEmailAddress(id *uuid.UUID) core.Status {
	if id == nil {
		return *core.UnprocessableContentError("ID is required")
	}

	result, err := p.db.Exec("delete from email_addresses where id = ? and generated = 0", id.String())
	if err != nil {
		return *core.InternalError(err.Error())
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return *core.NotFoundError("email address not found")
	}

	return *core.StatusNoContent()
}
