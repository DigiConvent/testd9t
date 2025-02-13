package api

import (
	sys_router "github.com/DigiConvent/testd9t/delivery/api/sys"
	services "github.com/DigiConvent/testd9t/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, s *services.Services) {
	sysRouter := sys_router.NewSysRouter(s.SysService, s.PostService)
	apiRoutes := router.Group("/api")

	apiRoutes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"sys": gin.H{
				"status": "/api/sys/status",
			},
		})
	})
	sysRoutes := apiRoutes.Group("/sys")

	sysRoutes.GET("/status", sysRouter.SystemStatusGet)
}
