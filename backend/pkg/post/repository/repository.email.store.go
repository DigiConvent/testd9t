package post_repository

import (
	"os"
	"path"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) StoreEmail(email *post_domain.EmailWrite) core.Status {
	if email == nil {
		return *core.UnprocessableContentError("email is required")
	}

	if email.To == "" {
		return *core.UnprocessableContentError("email to is required")
	}

	toId, status := p.GetEmailAddressByName(strings.Split(email.To, "@")[0])

	if status.Err() {
		return status
	}

	id, _ := uuid.NewV7()

	emailFolder := path.Join(os.Getenv(constants.DATABASE_PATH), "post", "email", id.String(), "attachments")
	err := os.MkdirAll(emailFolder, 0755)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	var notes []string
	for filename, attachment := range email.Attachments {
		err = os.WriteFile(path.Join(emailFolder, "attachments", filename), attachment, 0644)
		if err != nil {
			notes = append(notes, "Could not store attachment "+filename+": "+err.Error())
		}
	}

	err = os.WriteFile(path.Join(emailFolder, "body"), []byte(email.Body), 0644)
	if err != nil {
		notes = append(notes, "Could not store email body: "+err.Error())
	}

	_, err = p.db.Exec("insert into emails (id, from, to, subject, body, attachments) values (?, ?, ?, ?, ?, ?)",
		id.String(),
		email.From,
		toId.ID,
		email.Subject,
		email.Body,
		strings.Join(notes, "\n"))
	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
