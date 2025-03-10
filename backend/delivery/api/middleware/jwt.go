package api_middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"

	"github.com/DigiConvent/testd9t/core/log"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var pubkey *rsa.PublicKey = nil

func getPubKey() *rsa.PublicKey {
	if pubkey == nil {
		c, err := os.ReadFile(iam_setup.JwtPublicKeyPath())
		if err != nil {
			log.Error(err.Error())
		}

		block, _ := pem.Decode(c)
		if block == nil {
			log.Error("Could not decode public key")
			return nil
		}
		pPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			log.Error(err.Error())
		}

		pubkey = pPubKey.(*rsa.PublicKey)
	}
	return pubkey
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.Next()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return getPubKey(), nil
		})

		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set(ContextField, claims[ContextField])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
