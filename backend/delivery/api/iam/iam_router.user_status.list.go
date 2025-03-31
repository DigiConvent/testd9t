package iam_router

import (
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) ListUserStatus(ctx *gin.Context) {
	userStatus, status := router.iamService.ListUserStatuses()

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	} else {
		ctx.JSON(status.Code, gin.H{"items": userStatus})
	}
}
