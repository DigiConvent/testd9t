package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRead struct {
	ID           uuid.UUID  `json:"id"`
	Emailaddress string     `json:"emailaddress"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	Enabled      bool       `json:"enabled"`
}

type UserWrite struct {
	Emailaddress string    `json:"emailaddress" validate:"required,email"`
	FirstName    string    `json:"first_name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	DateOfBirth  time.Time `json:"date_of_birth" validate:"required"`
	UserStatus   uuid.UUID `json:"status_id"`
}

type UserSetEnabled struct {
	Enabled bool `json:"enabled"`
}

type UserFacade struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Implied    bool      `json:"implied"`
	StatusID   uuid.UUID `json:"status_id"`
	StatusName string    `json:"status_name"`
}

type UserProfile struct {
	User        *UserRead                `json:"user"`
	UserStatus  []*UserHasStatusRead     `json:"status"`
	Groups      []*PermissionGroupFacade `json:"groups"`
	Permissions []*PermissionFacade      `json:"permissions"`
}

type UserFilterSort struct {
	Filter struct {
		Emailaddress *string
		FirstName    *string
		LastName     *string
		DateOfBirth  *string
	}
	Sort struct {
		Field string
		Asc   bool
	}
	Page         int
	ItemsPerPage int
}
