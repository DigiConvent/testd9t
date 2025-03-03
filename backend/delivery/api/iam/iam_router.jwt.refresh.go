package iam_router

import (
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) RefreshJwt(ctx *gin.Context) {
	var telegramAuthentication TelegramAuthPayload
	err := ctx.BindJSON(&telegramAuthentication)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
	} else {

		ctx.JSON(200, gin.H{
			"token": "to be implemented",
		})
	}
}
