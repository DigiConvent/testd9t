package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) CreateUserStatus(ctx *gin.Context) {
	var userStatusWrite iam_domain.UserStatusWrite

	if err := ctx.ShouldBindJSON(&userStatusWrite); err != nil {
		ctx.JSON(422, gin.H{"error": err.Error()})
		return
	}
	id, status := router.iamService.CreateUserStatus(&userStatusWrite)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{"error": status.Message})
		return
	}
	ctx.JSON(status.Code, gin.H{
		"id": id,
	})
}
