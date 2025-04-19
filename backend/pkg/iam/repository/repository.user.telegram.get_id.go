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
)

func (r *IAMRepository) GetTelegramID(dataString, botToken string) (*int, core.Status) {
	if dataString == "" {
		return nil, *core.UnprocessableContentError("dataString cannot be empty")
	}
	if botToken == "" {
		return nil, *core.UnprocessableContentError("botToken cannot be empty")
	}

	query, _ := url.ParseQuery(dataString)

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

	secretKey := hmac.New(sha256.New, []byte("WebAppData"))
	secretKey.Write([]byte(botToken))

	computedHash := hmac.New(sha256.New, secretKey.Sum(nil))
	computedHash.Write([]byte(strings.Join(pairs, "\n")))

	signed := hex.EncodeToString(computedHash.Sum(nil))
	if signed != hash {
		return nil, *core.UnauthorizedError("signed is not the same as the hash")
	}

	userData := query.Get("user")
	var userObj struct {
		Id int `json:"id"`
	}
	err := json.Unmarshal([]byte(userData), &userObj)
	if err != nil {
		return nil, *core.UnprocessableContentError(err.Error())
	}

	return &userObj.Id, *core.StatusSuccess()
}
