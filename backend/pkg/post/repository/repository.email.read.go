package post_repository

import (
	"os"
	"path"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) ReadEmail(id *uuid.UUID) (*post_domain.EmailRead, core.Status) {
	if id == nil {
		return nil, *core.UnprocessableContentError("ID is required")
	}

	var email = &post_domain.EmailRead{}

	err := p.db.QueryRow("select id, from_email_address, to_email_address, subject, read_at, sent_at from emails where id = ?", id.String()).
		Scan(&email.ID, &email.Correspondent, &email.Mailbox, &email.Subject, &email.ReadAt, &email.SentAt)

	if err != nil {
		return nil, *core.NotFoundError("email not found")
	}

	var attachments []string
	attachmentsDir := path.Join(os.Getenv(constants.DATABASE_PATH), "post", "email", id.String(), "attachments")

	files, err := os.ReadDir(attachmentsDir)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	for _, file := range files {
		attachments = append(attachments, file.Name())
	}

	body, err := os.ReadFile(path.Join(os.Getenv(constants.DATABASE_PATH), "post", "email", id.String(), "contents"))
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}
	email.Body = string(body)
	email.Attachments = attachments

	return email, *core.StatusSuccess()
}
