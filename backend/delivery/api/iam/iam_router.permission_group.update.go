package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) UpdatePermissionGroup(ctx *gin.Context) {
	var permissionGroupWrite iam_domain.PermissionGroupWrite

	if err := ctx.ShouldBindJSON(&permissionGroupWrite); err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}
	status := router.iamService.UpdatePermissionGroup(router_utils.GetId(ctx), &permissionGroupWrite)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
