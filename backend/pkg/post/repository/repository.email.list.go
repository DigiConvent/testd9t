package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (p PostRepository) ListEmails(fs *post_domain.EmailFilterSort) (*pagination.Page[*post_domain.EmailFacade], core.Status) {
	if fs == nil {
		return nil, *core.UnprocessableContentError("Filter and sort is required")
	}

	recipientClause := ""
	if fs.Filter.Recipient != nil {
		recipientClause = "where to_email_address = '" + fs.Filter.Recipient.String() + "'"
	}

	emails := []*post_domain.EmailFacade{}
	rows, err := p.db.Query("select id, to_email_address, from, subject, sent_at from emails " + recipientClause)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	for rows.Next() {
		email := &post_domain.EmailFacade{}
		err := rows.Scan(&email.ID, &email.To, &email.From, &email.Subject, &email.SentAt)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}

		emails = append(emails, email)
	}

	var page = &pagination.Page[*post_domain.EmailFacade]{
		Items:        emails,
		Page:         fs.Page,
		ItemsPerPage: fs.ItemsPerPage,
	}

	err = p.db.QueryRow("select count(*) from emails " + recipientClause).Scan(&page.ItemsCount)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	return page, *core.StatusSuccess()
}
