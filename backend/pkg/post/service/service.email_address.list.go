package post_service

import (
	"github.com/DigiConvent/testd9t/core"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func (s PostService) ListEmailAddresses(address string) ([]post_domain.EmailAddressRead, *core.Status) {
	panic("unimplemented")
}
