package router

import (
	"sync"

	"github.com/DigiConvent/testd9t/core/log"
	"github.com/DigiConvent/testd9t/delivery/api"
	services "github.com/DigiConvent/testd9t/pkg"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(services *services.Services) {
	mainRouter := gin.Default()
	mainRouter.RedirectTrailingSlash = true

	api.RegisterRoutes(mainRouter, services)

	serveFrontend(mainRouter)

	if sys_domain.ProgramVersion == "dev" {
		log.Info("Development mode")
		mainRouter.Use(gin.Logger())
		mainRouter.Use(gin.Recovery())
		runHttp(mainRouter)
	} else {
		var waitGroup sync.WaitGroup
		runHttps(mainRouter)
		waitGroup.Wait()
	}
}
