package iam_domain

import (
	"time"

	"github.com/google/uuid"
)

type UserStatusProfile struct {
	UserStatus *UserStatusRead `json:"user_status"`
	Users      []*UserFacade   `json:"users"`
}

type UserStatusRead struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Abbr        string    `json:"abbr"`
	Description string    `json:"description"`
	Archived    bool      `json:"archived"`
}

type UserStatusWrite struct {
	Name        string     `json:"name"`
	Abbr        string     `json:"abbr"`
	Description string     `json:"description"`
	Archived    bool       `json:"archived"`
	Parent      *uuid.UUID `json:"parent"`
}

type AddUserStatusToUser struct {
	UserID   uuid.UUID `json:"user"`
	StatusID uuid.UUID `json:"user_status"`
	When     time.Time `json:"when"`
}

type UserHasStatusFacade struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Abbr string    `json:"abbr"`
}

type UserHasStatusRead struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Abbr        string    `json:"abbr"`
	Start       time.Time `json:"start"`
	Description string    `json:"description"`
}
