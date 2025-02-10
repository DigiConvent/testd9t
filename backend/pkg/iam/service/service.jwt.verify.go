package iam_service

import (
	"errors"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/log"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (service *IAMService) VerifyJwt(token string) (*uuid.UUID, *core.Status) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		log.Success("Verified JWT")
		return &service.IAMRepository.GetPrivateKey().PublicKey, nil
	})

	if err != nil {
		return nil, core.UnauthorizedError(err.Error())
	}

	id := parsedToken.Claims.(jwt.MapClaims)["id"]
	if id == nil {
		return nil, core.UnauthorizedError("invalid token")
	}
	userId := uuid.MustParse(id.(string))

	return &userId, core.StatusSuccess()
}
