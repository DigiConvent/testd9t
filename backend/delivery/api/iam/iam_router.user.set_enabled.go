package iam_router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (router *IamRouter) SetEnabledUser(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	parsedId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(422, gin.H{"error": "Invalid id"})
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

	fmt.Println("Setting enabled for user", id, "to", payload.Enabled)

	status := router.iamService.SetEnabled(&parsedId, payload.Enabled)
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
