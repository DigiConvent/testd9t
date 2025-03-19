package sys_router

import (
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/DigiConvent/testd9t/core/file_repo"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (r *SysRouter) RefreshInstallation(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Bearer token"})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := verifyOicdToken(tokenString)
	if err != nil {
		log.Println("Invalid GitHub OIDC token:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OIDC token"})
		return
	}

	if claims.Repository != expectedRepo {
		log.Println("Unauthorized repository:", claims.Repository)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized repository"})
		return
	}

	status := r.sysService.RefreshInstallation()
	if status.Err() {
		c.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Webhook received successfully"})
}

const expectedRepo = file_repo.GHUser + "/" + file_repo.GHRepo
const githubOicdIssuer = "https://token.actions.githubusercontent.com"

type GitHubClaims struct {
	jwt.StandardClaims
	Repository string `json:"repository"`
	Actor      string `json:"actor"`
}

func fetchGitHubKeys() (*rsa.PublicKey, error) {
	resp, err := http.Get(githubOicdIssuer + "/.well-known/jwks")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Keys []struct {
			Kid string `json:"kid"`
			N   string `json:"n"`
			E   string `json:"e"`
		} `json:"keys"`
	}

	if len(data.Keys) > 0 {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(data.Keys[0].N))
	}
	return nil, fmt.Errorf("no keys found")
}

func verifyOicdToken(tokenString string) (*GitHubClaims, error) {
	key, err := fetchGitHubKeys()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &GitHubClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*GitHubClaims); ok && token.Valid {

		if claims.Issuer != githubOicdIssuer {
			return nil, fmt.Errorf("invalid issuer")
		}

		if claims.Repository != expectedRepo {
			return nil, fmt.Errorf("unauthorized repository: %s", claims.Repository)
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
