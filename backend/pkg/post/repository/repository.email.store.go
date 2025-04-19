package post_repository

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

func (p PostRepository) StoreEmail(email *post_domain.EmailWrite) core.Status {
	if email == nil {
		return *core.UnprocessableContentError("email is required")
	}

	if email.Correspondent == "" {
		return *core.UnprocessableContentError("correspondent is required")
	}

	mailbox, status := p.GetEmailAddressByName(strings.Split(email.Mailbox, "@")[0])

	if status.Err() {
		return status
	}

	id, _ := uuid.NewV7()

	emailFolder := path.Join(os.Getenv(constants.DATABASE_PATH), "post", "email", id.String())
	attachmentsFolder := path.Join(emailFolder, "attachments")
	err := os.MkdirAll(attachmentsFolder, 0755)
	if err != nil {
		return *core.InternalError(err.Error())
	}

	var notes []string
	log.Info("Found " + strconv.Itoa(len(email.Attachments)) + " attachments")
	for filename, attachment := range email.Attachments {
		err = os.WriteFile(path.Join(attachmentsFolder, filename), attachment, 0644)
		if err != nil {
			notes = append(notes, "Could not store attachment "+filename+": "+err.Error())
		}
	}

	err = os.WriteFile(path.Join(emailFolder, "html"), []byte(email.Html), 0644)
	if err != nil {
		notes = append(notes, "Could not store email body: "+err.Error())
	}

	log.Info("Notes: " + strings.Join(notes, "\n"))
	_, err = p.db.Exec("insert into emails (id, correspondent, mailbox, subject, notes) values (?, ?, ?, ?, ?)",
		id.String(),
		email.Correspondent,
		mailbox.Id,
		email.Subject,
		email.Html,
		strings.Join(notes, "\n"))
	if err != nil {
		return *core.InternalError(err.Error())
	}

	return *core.StatusSuccess()
}
