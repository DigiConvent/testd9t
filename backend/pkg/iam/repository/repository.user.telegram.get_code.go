package iam_repository

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) GetTelegramRegistrationCode(userId *uuid.UUID) (string, core.Status) {
	if userId == nil {
		return "", *core.UnprocessableContentError("ID is required")
	}
	user := r.db.QueryRow(`select email from users where id = ?`, userId.String())

	var email string
	err := user.Scan(&email)
	if err != nil {
		return "", core.Status{Code: 500, Message: "Failed to find user with email"}
	}

	code := GenerateCode(userId.String(), time.Now(), 10*time.Minute)

	last := 0
	for {
		newTime := time.Now().Add(time.Duration(last+1) * time.Minute)
		expectedCode := GenerateCode(userId.String(), newTime, 10*time.Minute)
		if expectedCode != code {
			break
		}
		last = last + 1
	}

	return code, core.Status{
		Code:    200,
		Message: time.Now().Add(time.Duration(last) * time.Minute).Format(time.RFC3339),
	}
}
