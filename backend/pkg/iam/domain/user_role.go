package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRoleProfile struct {
	UserRole *UserRoleRead `json:"user_role"`
	Members  []*UserFacade `json:"members"`
}

type UserRoleRead struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Abbr        string    `json:"abbr"`
	Description string    `json:"description"`
}

type UserRoleWrite struct {
	Name        string     `json:"name"`
	Abbr        string     `json:"abbr"`
	Description string     `json:"description"`
	Parent      *uuid.UUID `json:"parent"`
}

type AddUserRoleToUser struct {
	UserID uuid.UUID `json:"user"`
	RoleID uuid.UUID `json:"user_role"`
	When   time.Time `json:"when"`
}

type UserHasRoleFacade struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Abbr string    `json:"abbr"`
}

type UserHasRoleRead struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Abbr        string    `json:"abbr"`
	Start       time.Time `json:"start"`
	Description string    `json:"description"`
}
