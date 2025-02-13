package router

import (
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	"github.com/gin-gonic/gin"
)

func runHttp(router *gin.Engine) {
	go func() {
		err := router.Run(":" + os.Getenv(constants.HTTP_PORT))
		if err != nil {
			panic("failed to start dev server: " + err.Error())
		} else {
			log.Info("Server started at http://" + os.Getenv(constants.DOMAIN) + ":" + os.Getenv(constants.HTTP_PORT))
		}
	}()
}
