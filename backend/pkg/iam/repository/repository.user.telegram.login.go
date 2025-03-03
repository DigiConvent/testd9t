// exempt from testing
package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetUserByTelegramID(id *int) (*uuid.UUID, core.Status) {
	userRow := r.db.QueryRow(`select id from users where telegram_id = ?`, id)

	var userId uuid.UUID

	err := userRow.Scan(&userId)
	if err != nil {
		return nil, *core.NotFoundError(err.Error())
	}
	return &userId, *core.StatusSuccess()
}
