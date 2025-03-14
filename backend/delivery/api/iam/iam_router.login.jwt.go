package iam_router

import (
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoginWithJwt implements IamRouterInterface.
func (router *IamRouter) LoginWithJwt(ctx *gin.Context) {
	rawId := ctx.GetString(api_middleware.ContextField)
	parsedId, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	token, status := router.iamService.GenerateJwt(&parsedId)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}
