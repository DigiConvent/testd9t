package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (p PostRepository) ListEmailAddresses() ([]post_domain.EmailAddressRead, core.Status) {
	results := []post_domain.EmailAddressRead{}
	rows, err := p.db.Query(`select id, name, domain from email_addresses`)

	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	for rows.Next() {
		result := post_domain.EmailAddressRead{}
		err := rows.Scan(&result.Id, &result.Name, &result.Domain)
		if err != nil {
			return nil, *core.InternalError(err.Error())
		}
		results = append(results, result)
	}

	return results, *core.StatusSuccess()
}
