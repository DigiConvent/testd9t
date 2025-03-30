package api_middleware

import (
	"net/http"
	"strings"

	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IamMiddleware struct {
	IamService iam_service.IAMServiceInterface
}

func (i *IamMiddleware) RequiresAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawUserId := c.GetString(ContextField)
		if rawUserId == "" {
			c.JSON(403, gin.H{"error": "Not logged in"})
			c.Abort()
			return
		}
		parsedUserId, err := uuid.Parse(rawUserId)

		if err != nil {
			c.JSON(422, gin.H{"error": "Invalid id"})
			c.Abort()
			return
		}
		// make sure that this user is enabled
		isEnabled, _ := i.IamService.IsEnabled(&parsedUserId)
		if !isEnabled {
			c.JSON(403, gin.H{"error": "User is disabled"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (i *IamMiddleware) RequiresPermission(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions = append(permissions, "super")
		rawUserId := c.GetString(ContextField)
		userId, err := uuid.Parse(rawUserId)

		if err != nil {
			c.JSON(403, gin.H{"error": "Not logged in"})
			c.Abort()
			return
		}

		for _, permission := range permissions {
			if strings.Contains(permission, ":") {
				segments := strings.Split(permission, ".")
				for i := 1; i < len(segments); i++ {
					param, e := strings.CutPrefix(segments[i], ":")
					if e {
						val := c.Param(param)
						segments[i] = val
					}
				}
				permission = strings.Join(segments, ".")
			}

			exists := i.IamService.UserHasPermission(&userId, permission)
			if exists {
				isEnabled, _ := i.IamService.IsEnabled(&userId)
				if !isEnabled {
					c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
					c.Abort()
					return
				}
				c.Set("permission", permission)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}

type IamMiddlewareInterface interface {
	RequiresPermission(permissions ...string) gin.HandlerFunc
	RequiresAuthentication() gin.HandlerFunc
}

func NewIamMiddleware(iamService iam_service.IAMServiceInterface) IamMiddlewareInterface {
	return &IamMiddleware{
		IamService: iamService,
	}
}
