package iam_service

import (
	"time"

	"github.com/DigiConvent/testd9t/core"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (service *IAMService) GenerateJwt(userId *uuid.UUID) (string, *core.Status) {
	if userId == nil {
		return "", core.UnprocessableContentError("ID is required")
	}
	privKey := service.repository.GetPrivateKey()

	user, _ := service.repository.GetUserByID(userId)

	telegramId, _ := service.repository.GetUserTelegramID(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"id":   userId.String(),
		"user": user,
		"tgid": telegramId,
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString(privKey)
	if err != nil {
		return "", core.BadRequestError(err.Error())
	}

	return tokenString, core.StatusSuccess()
}
