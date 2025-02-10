// exempt from testing
package iam_repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"sort"
	"strings"

	"github.com/DigiConvent/testd9t/core"
	uuid "github.com/google/uuid"
)

func (r *IAMRepository) VerifyTelegramUser(body string) (*uuid.UUID, core.Status) {
	if body == "" {
		return nil, *core.UnprocessableContentError("body cannot be empty")
	}
	row := r.db.QueryRow("select telegram_bot_token from config limit 1")
	var botToken string
	err := row.Scan(&botToken)
	if err != nil {
		return nil, *core.InternalError(err.Error())
	}

	query, _ := url.ParseQuery(body)

	var hash string
	var pairs []string = []string{}

	for key, val := range query {
		if key == "hash" {
			hash = val[0]
			continue
		} else {
			pairs = append(pairs, key+"="+val[0])
		}
	}

	sort.Strings(pairs)

	skHmac := hmac.New(sha256.New, []byte("WebAppData"))
	skHmac.Write([]byte(botToken))

	impHmac := hmac.New(sha256.New, skHmac.Sum(nil))
	impHmac.Write([]byte(strings.Join(pairs, "\n")))

	signed := hex.EncodeToString(impHmac.Sum(nil))
	if signed != hash {
		return nil, *core.UnauthorizedError("signed is not the same as the hash")
	}

	userData := query.Get("user")
	var userObj struct {
		ID int `json:"id"`
	}
	err = json.Unmarshal([]byte(userData), &userObj)
	if err != nil {
		return nil, *core.UnprocessableContentError(err.Error())
	}

	userId := userObj.ID

	userRow := r.db.QueryRow(`select id from users where telegram_id = ?`, userId)

	var uid uuid.UUID
	var str string

	err = userRow.Scan(&str)
	if err != nil {
		return nil, *core.NotFoundError(err.Error())
	}
	uid, err = uuid.Parse(str)
	if err != nil {
		return nil, *core.UnprocessableContentError(err.Error())
	}
	return &uid, *core.StatusSuccess()
}
