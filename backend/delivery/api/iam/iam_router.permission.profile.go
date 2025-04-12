package iam_router

import (
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) GetPermissionProfile(ctx *gin.Context) {
	name := ctx.Params.ByName("name")

	permission, status := router.iamService.GetPermissionProfile(name)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, permission)
	}
}
