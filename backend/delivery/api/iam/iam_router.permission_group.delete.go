package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) DeletePermissionGroup(ctx *gin.Context) {
	permissionGroupId := router_utils.GetId(ctx)

	status := router.iamService.DeletePermissionGroup(permissionGroupId, false)
	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, gin.H{})
	}
}
