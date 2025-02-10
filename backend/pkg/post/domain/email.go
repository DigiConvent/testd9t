package post_domain

import "github.com/google/uuid"

type EmailAddressWrite struct {
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
}

type EmailAddressRead struct {
	ID     *uuid.UUID `json:"id"`
	Name   string     `json:"name"`
	Domain string     `json:"domain"`
}
