package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetCredentials(email string, password string) (*uuid.UUID, core.Status) {
	result := r.db.QueryRow("select id from users where email = ? and password = ?", email, password)
	var id uuid.UUID
	err := result.Scan(&id)
	if err != nil {
		return nil, *core.NotFoundError("User not found")
	}
	return &id, *core.StatusSuccess()
}
