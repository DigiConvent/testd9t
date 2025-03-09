package iam_router

import "github.com/gin-gonic/gin"

func (router *IamRouter) ListPermissionGroups(ctx *gin.Context) {
	permissionGroups, status := router.iamService.ListPermissionGroups()

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	} else {
		ctx.JSON(status.Code, permissionGroups)
	}
}
