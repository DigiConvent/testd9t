package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) UpdateUser(ctx *gin.Context) {
	var userId *uuid.UUID
	if ctx.Request.URL.Path == "/api/iam/user/me/" {
		userId = router_utils.GetUserId(ctx)
	} else {
		userId = router_utils.GetId(ctx)
	}

	var user iam_domain.UserWrite
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}
	status := router.iamService.UpdateUser(userId, &user)
	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, gin.H{})
	}
}
