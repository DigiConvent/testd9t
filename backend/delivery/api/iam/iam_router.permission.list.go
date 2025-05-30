package iam_router

import "github.com/gin-gonic/gin"

func (router *IamRouter) ListPermissions(ctx *gin.Context) {
	permissions, status := router.iamService.ListPermissions()

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	} else {
		ctx.JSON(status.Code, permissions)
	}
}
