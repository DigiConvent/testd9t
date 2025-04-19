package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) RemoveUserFromUserRole(ctx *gin.Context) {
	var payload iam_domain.UserBecameRoleWrite

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := router.iamService.RemoveUserFromUserRole(&payload)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
