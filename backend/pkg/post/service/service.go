package post_service

import (
	"os"

	"github.com/DigiConvent/testd9t/core"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	"github.com/google/uuid"
)

type PostServiceInterface interface {
	CreateEmailAddress(credentials *post_domain.EmailAddressWrite) (*uuid.UUID, *core.Status)
	ReadEmailAddress(id *uuid.UUID) (*post_domain.EmailAddressRead, *core.Status)
	DeleteEmailAddress(id *uuid.UUID) *core.Status
	ListEmailAddresses() ([]post_domain.EmailAddressRead, *core.Status)
	UpdateEmailAddresses(id *uuid.UUID, credentials *post_domain.EmailAddressWrite) *core.Status

	SendEmail(from *uuid.UUID, to, subject, body string) *core.Status
}

type PostService struct {
	repository post_repository.PostRepositoryInterface
	address    string
}

func NewPostService(repository post_repository.PostRepositoryInterface, startSmtpServer bool) PostServiceInterface {
	postService := PostService{
		repository: repository,
		address:    ":" + os.Getenv(constants.SMTP_PORT),
	}

	if startSmtpServer {
		log.Info("Starting smtp server on " + postService.address)
		go postService.startSmtpServer()
	} else {
		log.Info("Skipping smtp server start")
	}
	return postService
}

func (s *PostService) StartSmtpServer() {
}
