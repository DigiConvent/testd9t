package api

import (
	sys_router "github.com/DigiConvent/testd9t/delivery/api/sys"
	services "github.com/DigiConvent/testd9t/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, s *services.Services) {
	sysRouter := sys_router.NewSysRouter(s.SysService)
	apiRoutes := router.Group("/api")

	sysRoutes := apiRoutes.Group("/sys")

	sysRoutes.GET("/status", sysRouter.SystemStatusGet)
}
