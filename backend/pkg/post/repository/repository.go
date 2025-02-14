package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/pagination"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

type PostRepositoryInterface interface {
	CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, core.Status)
	ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, core.Status)
	DeleteEmailAddress(id *uuid.UUID) core.Status
	ListEmailAddresses() ([]post_domain.EmailAddressRead, core.Status)
	UpdateEmailAddress(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) core.Status

	GetEmailAddressByName(name string) (*post_domain.EmailAddressRead, core.Status)

	StoreEmail(email *post_domain.EmailWrite) core.Status
	ReadEmail(id *uuid.UUID) (*post_domain.EmailRead, core.Status)
	ListEmails(fs *post_domain.EmailFilterSort) (*pagination.Page[*post_domain.EmailFacade], core.Status)
}

type PostRepository struct {
	db db.DatabaseInterface
}

func NewPostRepository(db db.DatabaseInterface) PostRepositoryInterface {
	return PostRepository{
		db: db,
	}
}
