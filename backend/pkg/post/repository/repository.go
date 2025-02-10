package post_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	"github.com/google/uuid"
)

type PostRepositoryInterface interface {
	CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, core.Status)
	ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, core.Status)
	DeleteEmailAddress(id *uuid.UUID) core.Status
	ListEmailAddress(address string) ([]post_domain.EmailAddressRead, core.Status)
	UpdateEmailAddress(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) core.Status
}

type PostRepository struct {
	db db.DatabaseInterface
}

func NewPostRepository(db db.DatabaseInterface) PostRepositoryInterface {
	return PostRepository{
		db: db,
	}
}
