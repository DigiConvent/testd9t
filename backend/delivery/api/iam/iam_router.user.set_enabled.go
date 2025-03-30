package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) SetEnabledUser(ctx *gin.Context) {
	id := router_utils.GetId(ctx)
	userId := router_utils.GetUserId(ctx)

	if id.String() == userId.String() {
		ctx.JSON(422, gin.H{
			"error": "Cannot change own status",
		})
		return
	}

	var payload struct {
		Enabled bool `json:"enabled"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := router.iamService.SetEnabled(id, payload.Enabled)
	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	} else {
		ctx.JSON(status.Code, gin.H{
			"id": id,
		})
	}
}
