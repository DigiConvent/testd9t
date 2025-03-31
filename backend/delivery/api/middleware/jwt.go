package api_middleware

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"

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
		tokenString := c.GetHeader("Authentication")

		// custom auth method has no "Bearer " prefix but github does so we can't extract the user from this
		if tokenString == "" || strings.HasPrefix(tokenString, "Bearer ") {
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
			error := ""
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					error = fmt.Sprintln("Token has expired")
				} else {
					error = fmt.Sprintln("Token validation error:", err)
				}
			} else {
				error = fmt.Sprintln("Failed to parse token:", err)
			}
			log.Error(error)
			c.JSON(
				http.StatusUnauthorized,
				gin.H{"error": error},
			)
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
