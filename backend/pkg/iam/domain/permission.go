package iam_domain

import "github.com/google/uuid"

type PermissionWrite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Meta        string `json:"meta"`
}

type PermissionRead struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Meta        string    `json:"meta"`
}

type PermissionFacade struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Implied     bool      `json:"implied"`
}
