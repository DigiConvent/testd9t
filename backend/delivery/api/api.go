package api

import (
	iam_router "github.com/DigiConvent/testd9t/delivery/api/iam"
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	sys_router "github.com/DigiConvent/testd9t/delivery/api/sys"
	services "github.com/DigiConvent/testd9t/pkg"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, s *services.Services) {
	router.Use(api_middleware.JWTAuthMiddleware())
	iam := api_middleware.NewIamMiddleware(s.IamService)

	sysRouter := sys_router.NewSysRouter(s.SysService)
	iamRouter := iam_router.NewIamRouter(s.IamService, s.SysService)

	apiRoutes := router.Group("/api")

	iamRoutes := apiRoutes.Group("/iam")
	{
		jwtRoutes := iamRoutes.Group("/jwt")
		{
			jwtRoutes.POST("/refresh", iamRouter.LoginWithJwt)

		}
		loginRoutes := iamRoutes.Group("/login")
		{
			loginRoutes.POST("/credentials", iamRouter.LoginWithCredentials)
			loginRoutes.POST("/telegram", iamRouter.LoginWithTelegram)
			loginRoutes.POST("/telegram/connect", iamRouter.ConnectTelegramUser)
		}

		userRoutes := iamRoutes.Group("/user")
		{
			userRoutes.GET("/", iam.RequiresPermission("iam.user.list"), iamRouter.ListUsers)
			userRoutes.POST("/", iam.RequiresPermission("iam.user.create"), iamRouter.CreateUser)
			userRoutes.GET("/:id", iam.RequiresPermission("iam.user.get", "iam.user.get.:id"), iamRouter.GetUser)

			userRoutes.GET("/permission", iamRouter.ListPermissionsUser)
			userRoutes.GET("/permission/:id", iam.RequiresPermission("iam.user.get", "iam.user.get.:id"), iamRouter.ListPermissionsUser)
		}

		userStatusRoutes := iamRoutes.Group("/user-status")
		{
			userStatusRoutes.GET("/", iam.RequiresPermission("iam.user-status.list"), iamRouter.ListUserStatus)
			userStatusRoutes.POST("/", iam.RequiresPermission("iam.user-status.create"), iamRouter.CreateUserStatus)
			// userStatusRoutes.GET("/:id", iam.RequiresPermission("iam.user-status.get"), iamRouter.GetUserStatus)
		}

		permissionRoutes := iamRoutes.Group("/permission")
		{
			permissionRoutes.GET("/", iam.RequiresPermission("iam.permission.list"), iamRouter.ListPermissions)
		}

		permissionGroupRoutes := iamRoutes.Group("/permission-group")
		{
			permissionGroupRoutes.GET("/", iam.RequiresPermission("iam.permission-group.list"), iamRouter.ListPermissionGroups)
			permissionGroupRoutes.GET("/:id", iam.RequiresPermission("iam.permission-group.get", "iam.permission-group.get.:id"), iamRouter.GetPermissionGroup)

			permissionGroupRoutes.POST("/", iam.RequiresPermission("iam.permission-group.create"), iamRouter.CreatePermissionGroup)
			permissionGroupRoutes.POST("/:id", iam.RequiresPermission("iam.permission-group.update"), iamRouter.UpdatePermissionGroup)
			permissionGroupRoutes.GET("/profile/:id", iam.RequiresPermission("iam.permission-group.get", "iam.permission-group.get.:id"), iamRouter.ProfilePermissionGroup)
		}
	}

	sysRoutes := apiRoutes.Group("/sys")
	{
		sysRoutes.GET("/status", iam.RequiresPermission("sys"), sysRouter.GetStatus)
		sysRoutes.GET("/installation/refresh", sysRouter.RefreshInstallation)
		sysRoutes.POST("/logo/small", iam.RequiresPermission("sys"), sysRouter.SetSmallLogo)
		sysRoutes.POST("/logo/large", iam.RequiresPermission("sys"), sysRouter.SetLargeLogo)
	}
}
