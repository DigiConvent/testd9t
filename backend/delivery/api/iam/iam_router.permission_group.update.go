package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) UpdatePermissionGroup(ctx *gin.Context) {
	var permissionGroupWrite iam_domain.PermissionGroupWrite
	id := ctx.Params.ByName("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	if err := ctx.ShouldBindJSON(&permissionGroupWrite); err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}
	status := router.iamService.UpdatePermissionGroup(&parsedId, &permissionGroupWrite)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
