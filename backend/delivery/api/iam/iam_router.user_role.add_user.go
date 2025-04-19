package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) AddUserToUserRole(ctx *gin.Context) {
	var addUserToUserRole iam_domain.UserBecameRoleWrite

	if err := ctx.ShouldBindJSON(&addUserToUserRole); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := router.iamService.AddUserToUserRole(&addUserToUserRole)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
