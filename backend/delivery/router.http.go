package router

import "github.com/gin-gonic/gin"

func runHttp(router *gin.Engine) {
	go func() {
		err := router.Run(":80")
		if err != nil {
			panic("failed to start dev server: " + err.Error())
		}
	}()
}
