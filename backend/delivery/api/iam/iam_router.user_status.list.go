package iam_router

import (
	"fmt"

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
		fmt.Println(userStatus)
		ctx.JSON(status.Code, gin.H{"items": userStatus})
	}
}
