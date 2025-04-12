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
	IsGroup     bool     `json:"is_group"`
	IsNode      bool     `json:"is_node"`
	Permissions []string `json:"permissions"`
}

type PermissionGroupSetParent struct {
	ID     *uuid.UUID `json:"id"`
	Parent *uuid.UUID `json:"parent"`
}

type PermissionGroupRead struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Abbr        string              `json:"abbr"`
	Description string              `json:"description"`
	Parent      *uuid.UUID          `json:"parent"`
	Meta        string              `json:"meta"`
	IsGroup     bool                `json:"is_group"`
	IsNode      bool                `json:"is_node"`
	Generated   bool                `json:"generated"`
	Permissions []*PermissionFacade `json:"permissions"`
}

type PermissionGroupFacade struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Abbr      string     `json:"abbr"`
	IsGroup   bool       `json:"is_group"`
	IsNode    bool       `json:"is_node"`
	Meta      *string    `json:"meta"`
	Implied   bool       `json:"implied"`
	Parent    *uuid.UUID `json:"parent"`
	Generated bool       `json:"generated"`
}

type PermissionGroupProfile struct {
	PermissionGroup *PermissionGroupRead     `json:"permission_group"`
	Permissions     []*PermissionFacade      `json:"permissions"`
	Members         []*UserFacade            `json:"members"`
	Ancestors       []*PermissionGroupFacade `json:"ancestors"`
	Descendants     []*PermissionGroupFacade `json:"descendants"`
}

type AddUserToPermissionGroupWrite struct {
	PermissionGroup *uuid.UUID `json:"permission_group"`
	User            *uuid.UUID `json:"user"`
	Start           *time.Time `json:"start"`
	End             *time.Time `json:"end"`
}
