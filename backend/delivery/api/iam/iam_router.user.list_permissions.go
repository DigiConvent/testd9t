package iam_router

import (
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) ListPermissionsUser(ctx *gin.Context) {
	var rawId string
	permission := ctx.GetString("permission")

	if ctx.Param("id") != "" && (permission == "admin" || permission == "iam.user.get" || permission == "iam.user.get."+ctx.Param("id")) {
		rawId = ctx.Param("id")
	} else {
		id := ctx.GetString(api_middleware.ContextField)
		if id == "" {
			ctx.JSON(403, gin.H{"error": "Not logged in"})
			return
		}
		rawId = id
	}

	parsedId, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	permissions, status := router.iamService.ListUserPermissions(&parsedId)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	} else {
		ctx.JSON(status.Code, permissions)
	}
}
