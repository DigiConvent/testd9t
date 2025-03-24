package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) CreatePermissionGroup(ctx *gin.Context) {
	var permissionGroupWrite iam_domain.PermissionGroupWrite

	if err := ctx.ShouldBindJSON(&permissionGroupWrite); err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}
	id, status := router.iamService.CreatePermissionGroup(&permissionGroupWrite)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	ctx.JSON(status.Code, gin.H{
		"id": id,
	})
}
