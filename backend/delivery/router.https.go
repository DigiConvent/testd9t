package router

import (
	"net/http"
	"os"
	"sync"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/log"
	sys_setup "github.com/DigiConvent/testd9t/pkg/sys/setup"
	"github.com/gin-gonic/gin"
)

func runHttps(router *gin.Engine) {
	var waitGroup sync.WaitGroup
	runHttp2Https(&waitGroup)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		err := router.RunTLS(":"+os.Getenv(constants.HTTPS_PORT), sys_setup.TlsPublicKeyPath(), sys_setup.TlsPrivateKeyPath())
		if err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()
	waitGroup.Wait()
}

func runHttp2Https(waitGroup *sync.WaitGroup) {
	go func() {
		defer waitGroup.Done()
		subRouter := gin.New()
		subRouter.Use(func(ctx *gin.Context) {
			log.Info("Redirecting http to https://" + ctx.Request.Host + ctx.Request.RequestURI)
			ctx.Redirect(http.StatusMovedPermanently, "https://"+ctx.Request.Host+ctx.Request.RequestURI)
		})
		if err := subRouter.Run(":" + os.Getenv(constants.HTTP_PORT)); err != nil {
			log.Error("Could not start http redirect server: " + err.Error())
		}
	}()
}
