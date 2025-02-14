package post_domain

import (
	"time"

	"github.com/google/uuid"
)

type EmailRead struct {
	From        string     `json:"from"`
	To          *uuid.UUID `json:"to"`
	ID          *uuid.UUID `json:"id"`
	Subject     string     `json:"subject"`
	Body        string     `json:"body"`
	Attachments []string   `json:"attachments"`
	ReadAt      *time.Time `json:"read_at"`
	SentAt      *time.Time `json:"when"`
	Notes       string     `json:"notes"`
}

type EmailFacade struct {
	ID          uuid.UUID  `json:"id"`
	From        string     `json:"from"`
	To          string     `json:"to"`
	Subject     string     `json:"subject"`
	SentAt      *time.Time `json:"when"`
	ReadAt      *time.Time `json:"read_at"`
	Attachments []string   `json:"attachments"`
}

type EmailWrite struct {
	From        string            `json:"from"`
	To          string            `json:"to"`
	Subject     string            `json:"subject"`
	Body        string            `json:"body"`
	Attachments map[string][]byte `json:"attachments"`
	Notes       []string          `json:"notes"`
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
