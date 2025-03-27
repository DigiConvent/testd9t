package iam_router

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core/log"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) CreateUser(ctx *gin.Context) {
	var createUser iam_domain.UserWrite
	if err := ctx.ShouldBindJSON(&createUser); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(createUser)

	id, status := router.iamService.CreateUser(&createUser)
	if status != nil && status.Err() {
		log.Info(status.Message)
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
