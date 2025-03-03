package iam_router

import (
	"net/http"
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/gin-gonic/gin"
)

type TelegramAuthPayload struct {
	Payload string `json:"payload"`
}

func (router *IamRouter) LoginWithTelegram(ctx *gin.Context) {
	var telegramAuthentication TelegramAuthPayload
	err := ctx.BindJSON(&telegramAuthentication)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
	} else {
		id, status := router.iamService.LoginTelegramUser(telegramAuthentication.Payload, os.Getenv(constants.TELEGRAM_BOT_TOKEN))
		if status != nil && status.Err() {
			ctx.JSON(status.Code, gin.H{
				"status": status,
			})
			return
		}

		jwt, status := router.iamService.GenerateJwt(id)
		if status != nil && status.Err() {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": status.Message})
			return
		}

		ctx.JSON(status.Code, gin.H{
			"jwt": jwt,
		})
	}
}
