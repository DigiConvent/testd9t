package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type UserRoleProfile struct {
	*PermissionGroupProfile
	UserRole *UserRoleRead         `json:"user_role"`
	History  []*UserBecameRoleRead `json:"history"`
}

type UserRoleRead struct {
	Id          uuid.UUID `json:"id"`
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

type UserBecameRoleWrite struct {
	User    uuid.UUID `json:"user"`
	Role    uuid.UUID `json:"user_role"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Comment string    `json:"comment"`
}

type UserHasRoleFacade struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Abbr string    `json:"abbr"`
}

type UserRoleFacade struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserHasRoleRead struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Abbr    string    `json:"abbr"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Comment string    `json:"comment"`
}

type UserBecameRoleRead struct {
	Role  uuid.UUID  `json:"role"`
	User  UserFacade `json:"user"`
	Start time.Time  `json:"start"`
	End   time.Time  `json:"end"`
}
