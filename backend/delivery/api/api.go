package api

import (
	iam_router "github.com/DigiConvent/testd9t/delivery/api/iam"
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	sys_router "github.com/DigiConvent/testd9t/delivery/api/sys"
	services "github.com/DigiConvent/testd9t/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, s *services.Services) {
	iamRouter := iam_router.NewIamRouter(s.IamService)
	iamRoutes := router.Group("/iam")
	iamRoutes.POST("/login/credentials", iamRouter.LoginWithCredentials)
	iamRoutes.POST("/login/telegram", iamRouter.LoginWithTelegram)
	iamRoutes.POST("/jwt/refresh", iamRouter.)

	sys_router.SetupSysRoutes(router, s.SysService, s.PostService)
	apiRoutes := router.Group("/api")

	apiRoutes.Use(api_middleware.JWTAuthMiddleware())

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
