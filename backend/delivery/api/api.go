package api

import (
	sys_router "github.com/DigiConvent/testd9t/delivery/api/sys"
	services "github.com/DigiConvent/testd9t/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, s *services.Services) {
	sysRouter := sys_router.NewSysRouter(s.SysService)

	router.Group("/api")
	{
		router.Group("/sys")
		{
			router.GET("/status", sysRouter.SystemStatusGet)
		}
	}
}
