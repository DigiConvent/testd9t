package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type PermissionGroupWrite struct {
	Name        string   `json:"name"`
	Abbr        string   `json:"abbr"`
	Description string   `json:"description"`
	Parent      string   `json:"parent"`
	Permissions []string `json:"permissions"`
}

type PermissionGroupSetParent struct {
	Id     *uuid.UUID `json:"id"`
	Parent *uuid.UUID `json:"parent"`
}

type PermissionGroupRead struct {
	Id          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Abbr        string              `json:"abbr"`
	Description string              `json:"description"`
	Parent      *uuid.UUID          `json:"parent"`
	Meta        string              `json:"meta"`
	Generated   bool                `json:"generated"`
	Permissions []*PermissionFacade `json:"permissions"`
}

type PermissionGroupFacade struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Abbr      string     `json:"abbr"`
	Meta      *string    `json:"meta"`
	Implied   bool       `json:"implied"`
	Parent    *uuid.UUID `json:"parent"`
	Generated bool       `json:"generated"`
}

type PermissionGroupProfile struct {
	PermissionGroup *PermissionGroupRead     `json:"permission_group"`
	Permissions     []*PermissionFacade      `json:"permissions"`
	Users           []*UserFacade            `json:"users"`
	Ancestors       []*PermissionGroupFacade `json:"ancestors"`
	Descendants     []*PermissionGroupFacade `json:"descendants"`
}

type AddUserToPermissionGroupWrite struct {
	PermissionGroup *uuid.UUID `json:"permission_group"`
	User            *uuid.UUID `json:"user"`
	Start           *time.Time `json:"start"`
	End             *time.Time `json:"end"`
}
