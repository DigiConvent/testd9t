package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetUserTelegramID(id *uuid.UUID) (*int, core.Status) {
	result := r.db.QueryRow("select telegram_id from users where id = ?", id.String())
	var telegramID int
	err := result.Scan(&telegramID)

	if err != nil {
		return nil, *core.NotFoundError("User not found")
	}

	return &telegramID, *core.StatusSuccess()
}
