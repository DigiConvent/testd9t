package post_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (p PostRepository) GetEmailAddressByName(name string) (*post_domain.EmailAddressRead, core.Status) {
	result := &post_domain.EmailAddressRead{}
	err := p.db.QueryRow("select id, name, domain from email_addresses where name = ?", strings.ToLower(name)).Scan(&result.Id, &result.Name, &result.Domain)
	if err != nil {
		return nil, *core.NotFoundError("email address not found")
	}
	return result, *core.StatusSuccess()
}
