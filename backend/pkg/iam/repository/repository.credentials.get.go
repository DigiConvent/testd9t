package iam_repository

import (
	"strings"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetCredentials(emailaddress string) (*uuid.UUID, string, core.Status) {
	result := r.db.QueryRow("select id, password from users where emailaddress = ?", strings.ToLower(emailaddress))
	var id uuid.UUID
	var hashedPassword string
	err := result.Scan(&id, &hashedPassword)
	if err != nil {
		return nil, "", *core.NotFoundError("User not found")
	}
	return &id, hashedPassword, *core.StatusSuccess()
}
