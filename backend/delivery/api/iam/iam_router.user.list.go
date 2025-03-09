package iam_router

import (
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) ListUsers(ctx *gin.Context) {
	var fs iam_domain.UserFilterSort

	if err := ctx.ShouldBindQuery(&fs); err != nil {
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
		return
	}

	users, status := router.iamService.ListUsers(&fs)
	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"status": status,
		})
		return
	}

	ctx.JSON(status.Code, users)
}
