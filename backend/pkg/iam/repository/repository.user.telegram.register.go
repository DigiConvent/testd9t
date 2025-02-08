package iam_repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

// imagine if a code is only valid for 10 minutes
// the code is generated at 10:09:56 so you have 4 seconds left to use it
// find the previous round minute so minutes/10 * 10
// 9/10 * 10 = 0
// 10/10 * 10 = 10
// 11/10 * 10 = 10
// 19/10 * 10 = 10
// 20/10 * 10 = 20
// 21/10 * 10 = 20
// 49/10 * 10 = 40
// 50/10 * 10 = 50
// 51/10 * 10 = 50
// 59/10 * 10 = 50
// 60/10 * 10 = 60
// 0/10 * 10 = 0

func previousRoundMinute(t time.Time) time.Time {
	minutes := t.Minute()
	roundedMinutes := 10 * (minutes / 10)
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), roundedMinutes, 0, 0, t.Location())
}

func GenerateCode(uuid string, t time.Time, validityPeriod time.Duration) string {
	previousRoundMinute := previousRoundMinute(t)
	expirationTime := previousRoundMinute.Add(validityPeriod).Unix()
	message := fmt.Sprintf("%d:%d", previousRoundMinute.Unix(), expirationTime)

	h := hmac.New(sha256.New, []byte(uuid))
	h.Write([]byte(message))
	code := hex.EncodeToString(h.Sum(nil))

	return code
}

func VerifyCode(uuid string, t time.Time, validityPeriod time.Duration, code string) bool {
	for i := -1; i < 2; i++ {
		newTime := t.Add(time.Duration(i*5) * time.Minute)
		expectedCode := GenerateCode(uuid, newTime, validityPeriod)
		if expectedCode == code {
			return true
		}
	}
	return false
}

func (r *IAMRepository) RegisterTelegramUser(telegramId int, email string, code string) core.Status {
	email = strings.ToLower(email)
	row := r.DB.QueryRow("select id from users where email = ?", email)

	var userId uuid.UUID
	err := row.Scan(&userId)
	if err != nil {
		return *core.InternalError("Failed to find user with email")
	}

	if !VerifyCode(userId.String(), time.Now(), 10*time.Minute, code) {
		return *core.InternalError("Invalid code")
	}

	result, err := r.DB.Exec("update users set telegram_id = ? where id = ?", telegramId, userId.String())
	if err != nil {
		return *core.InternalError("Failed to update user")
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return *core.InternalError("Failed to update user")
	}

	return *core.StatusSuccess()

}
