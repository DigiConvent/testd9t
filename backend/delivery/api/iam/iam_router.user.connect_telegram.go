package iam_router

import (
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) ConnectTelegramUser(ctx *gin.Context) {
	rawId := ctx.GetString(api_middleware.ContextField)
	if rawId == "" {
		ctx.JSON(403, gin.H{"error": "Not logged in"})
		return
	}
	parsedId, err := uuid.Parse(rawId)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	var telegramAuthentication TelegramAuthPayload
	err = ctx.BindJSON(&telegramAuthentication)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
	} else {
		status := router.iamService.ConnectTelegramUser(telegramAuthentication.Payload, os.Getenv(constants.TELEGRAM_BOT_TOKEN), &parsedId)

		if status.Err() {
			ctx.JSON(status.Code, gin.H{
				"error": status.Message,
			})
			return
		} else {
			ctx.JSON(status.Code, gin.H{
				"status": status.Message,
			})
		}
	}
}
