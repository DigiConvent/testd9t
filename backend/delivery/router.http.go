package router

import (
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/gin-gonic/gin"
)

func runHttp(router *gin.Engine) {
	go func() {
		err := router.Run(":" + os.Getenv(constants.HTTP_PORT))
		if err != nil {
			panic("failed to start dev server: " + err.Error())
		}
	}()
}
