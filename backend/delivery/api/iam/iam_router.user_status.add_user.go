package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) AddUserStatusUser(ctx *gin.Context) {
	var addUserToUserStatus iam_domain.UserBecameStatusWrite
	if err := ctx.ShouldBindJSON(&addUserToUserStatus); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := router.iamService.AddUserToUserStatus(&addUserToUserStatus)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}
	ctx.JSON(status.Code, gin.H{})
}
