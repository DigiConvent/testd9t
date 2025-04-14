package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	"github.com/gin-gonic/gin"
)

type EditPermissionGroupPermissions struct {
	Add    string `json:"add"`
	Remove string `json:"remove"`
}

func (router *IamRouter) PermissionGroupEditPermissions(ctx *gin.Context) {
	permissionGroupId := router_utils.GetId(ctx)

	var payload EditPermissionGroupPermissions
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(payload.Add) == 0 && len(payload.Remove) == 0 {
		ctx.JSON(422, gin.H{
			"error": "No changes",
		})
		return
	}

	if payload.Add != "" {
		status := router.iamService.AddPermissionToPermissionGroup(permissionGroupId, payload.Add)
		if status != nil && status.Err() {
			ctx.JSON(status.Code, gin.H{
				"error": status.Message,
			})
			return
		}
	}

	if payload.Remove != "" {
		status := router.iamService.RemovePermissionFromPermissionGroup(permissionGroupId, payload.Remove)
		if status != nil && status.Err() {
			ctx.JSON(status.Code, gin.H{
				"error": status.Message,
			})
			return
		}
	}

	ctx.JSON(204, gin.H{})
}
