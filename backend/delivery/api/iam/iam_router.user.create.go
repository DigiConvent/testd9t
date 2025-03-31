package iam_router

import (
	"bytes"
	"fmt"
	"io"

	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	"github.com/gin-gonic/gin"
)

func (router *IamRouter) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	var createUser iam_domain.UserWrite
	if err := ctx.ShouldBindJSON(&createUser); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, status := router.iamService.CreateUser(&createUser)
	if status != nil && status.Err() {
		fmt.Println(status.Message)
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}

	// rewind
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	var addUserToUserStatus iam_domain.AddUserStatusToUser
	if err := ctx.ShouldBindJSON(&addUserToUserStatus); err != nil {
		ctx.JSON(422, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("Created user but now it is time to add user status")
	addUserToUserStatus.UserID = *id
	fmt.Println(addUserToUserStatus.UserID)
	fmt.Println(addUserToUserStatus.StatusID)
	fmt.Println(addUserToUserStatus.When)

	status = router.iamService.AddUserStatus(&addUserToUserStatus)
	if status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
		return
	}
	ctx.JSON(status.Code, gin.H{
		"id": id,
	})
}
