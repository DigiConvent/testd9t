package iam_domain

import "github.com/google/uuid"

type PermissionGroupWrite struct {
	Name        string       `json:"name" validate:"required"`
	Abbr        string       `json:"abbr"`
	Description string       `json:"description"`
	Parent      string       `json:"parent"`
	IsGroup     bool         `json:"is_group"`
	IsNode      bool         `json:"is_node"`
	Permissions []*uuid.UUID `json:"permissions"`
}

type PermissionGroupSetParent struct {
	ID     uuid.UUID  `json:"id" validate:"required"`
	Parent *uuid.UUID `json:"parent" validate:"required"`
}

type PermissionGroupRead struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Abbr        string     `json:"abbr"`
	Description string     `json:"description"`
	Parent      *uuid.UUID `json:"parent"`
	IsGroup     bool       `json:"is_group"`
	IsNode      bool       `json:"is_node"`
	Generated   bool       `json:"generated"`
}

type PermissionGroupFacade struct {
	ID      uuid.UUID  `json:"id"`
	Name    string     `json:"name"`
	Abbr    string     `json:"abbr"`
	IsGroup bool       `json:"is_group"`
	Implied bool       `json:"implied"`
	Parent  *uuid.UUID `json:"parent"`
}

type PermissionGroupProfile struct {
	PermissionGroup *PermissionGroupRead `json:"permission_group"`
	Permissions     []*PermissionFacade  `json:"permissions"`
	Members         []*UserFacade        `json:"members"`
}
