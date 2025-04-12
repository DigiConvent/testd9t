package iam_domain

import (
	"github.com/google/uuid"
)

type UserRead struct {
	ID           uuid.UUID `json:"id"`
	Emailaddress string    `json:"emailaddress"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Enabled      bool      `json:"enabled"`
}

type UserWrite struct {
	Emailaddress string         `json:"emailaddress"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Extra        map[string]any `json:"-"`
}

type UserSetEnabled struct {
	Enabled bool `json:"enabled"`
}

type UserFacade struct {
	ID         uuid.UUID         `json:"id"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	StatusID   *uuid.UUID        `json:"status_id"`
	StatusName *string           `json:"status_name"`
	Roles      []*UserRoleFacade `json:"roles"`
	Implied    bool              `json:"implied"`
}

type UserProfile struct {
	User        *UserRead                `json:"user"`
	UserStatus  []*UserHasStatusRead     `json:"status"`
	UserRole    []*UserHasRoleRead       `json:"role"`
	Groups      []*PermissionGroupFacade `json:"groups"`
	Permissions []*PermissionFacade      `json:"permissions"`
}

type UserFilterSort struct {
	Filter struct {
		Emailaddress *string
		FirstName    *string
		LastName     *string
	}
	Sort struct {
		Field string
		Asc   bool
	}
	Page         int
	ItemsPerPage int
}
