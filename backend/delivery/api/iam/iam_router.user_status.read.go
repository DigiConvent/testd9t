package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) GetUserStatus(ctx *gin.Context) {
	userStatus, status := router.iamService.GetUserStatus(router_utils.GetId(ctx))

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, userStatus)
	}
}
