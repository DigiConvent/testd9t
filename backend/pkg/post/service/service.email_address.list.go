package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (s PostService) ListEmailAddresses() ([]post_domain.EmailAddressRead, *core.Status) {
	addresses, status := s.repository.ListEmailAddresses()

	return addresses, &status
}
