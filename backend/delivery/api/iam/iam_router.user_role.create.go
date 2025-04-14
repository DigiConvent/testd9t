package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) CreateUserRole(ctx *gin.Context) {
	var payload iam_domain.UserRoleWrite

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, status := router.iamService.CreateUserRole(&payload)
	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}

	ctx.JSON(status.Code, gin.H{
		"id": id,
	})
}
