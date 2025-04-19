package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) GetUserRoleProfile(ctx *gin.Context) {
	id := router_utils.GetId(ctx)
	userRole, status := router.iamService.GetUserRoleProfile(id)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, userRole)
	}
}
