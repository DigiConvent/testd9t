package iam_router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) GetUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
		return
	}

	userProfile, status := router.iamService.GetUser(&parsedId)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, userProfile)
	}
}
