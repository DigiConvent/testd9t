package iam_router

import (
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SetPasswordUserRequest struct {
	Password string `json:"password"`
}

func (router *IamRouter) SetPasswordUser(ctx *gin.Context) {
	var setPasswordRequest SetPasswordUserRequest

	var rawId string
	permission := ctx.GetString("permission")
	if ctx.Param("id") != "" && (permission == "super" || permission == "iam.user.set_password") {
		rawId = ctx.Param("id")
	} else {
		id := ctx.GetString(api_middleware.ContextField)
		if id == "" {
			ctx.JSON(403, gin.H{"error": "Not logged in"})
			return
		}
		rawId = id
	}

	if err := ctx.ShouldBindJSON(&setPasswordRequest); err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}

	parsedId, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	status := router.iamService.SetUserPassword(&parsedId, setPasswordRequest.Password)

	if status.Err() {
		ctx.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
