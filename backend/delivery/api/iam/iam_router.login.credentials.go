package iam_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Emailaddress string `json:"emailaddress"`
	Password     string `json:"password"`
}

func (r *IamRouter) LoginWithCredentials(ctx *gin.Context) {
	var credentials Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(400, gin.H{
			"status": err.Error(),
		})
	} else {
		id, status := r.iamService.LoginUser(credentials.Emailaddress, credentials.Password)
		if status != nil && status.Err() {
			ctx.JSON(status.Code, gin.H{
				"status": status,
			})
			return
		}

		jwt, status := r.iamService.GenerateJwt(id)
		if status != nil && status.Err() {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": status.Message})
			return
		}

		ctx.JSON(status.Code, gin.H{
			"jwt": jwt,
		})
	}
}
