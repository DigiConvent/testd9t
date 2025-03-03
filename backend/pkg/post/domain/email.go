package post_domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailRead struct {
	Mailbox       *uuid.UUID `json:"mailbox"`
	Correspondent string     `json:"correspondent"`
	ID            *uuid.UUID `json:"id"`
	Subject       string     `json:"subject"`
	Body          string     `json:"body"`
	Attachments   []string   `json:"attachments"`
	ReadAt        *time.Time `json:"read_at"`
	SentAt        *time.Time `json:"when"`
	Notes         string     `json:"notes"`
}

type EmailFacade struct {
	ID            *uuid.UUID `json:"id"`
	Mailbox       *uuid.UUID `json:"mailbox"`
	Correspondent string     `json:"correspondent"`
	Subject       string     `json:"subject"`
	SentAt        *time.Time `json:"when"`
	ReadAt        *time.Time `json:"read_at"`
	Attachments   []string   `json:"attachments"`
}

type EmailWrite struct {
	Mailbox       string            `json:"mailbox"`
	Correspondent string            `json:"correspondent"`
	Subject       string            `json:"subject"`
	Html          string            `json:"html"`
	Plain         string            `json:"plain"`
	Attachments   map[string][]byte `json:"attachments"`
}

type EmailFilterSort struct {
	Filter struct {
		Recipient *uuid.UUID
	}
	Sort struct {
		Field string
		Asc   bool
	}
	Page         int
	ItemsPerPage int
}
