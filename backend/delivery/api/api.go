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
			jwtRoutes.POST("/refresh/", iamRouter.LoginWithJwt)

		}
		loginRoutes := iamRoutes.Group("/login")
		{
			loginRoutes.POST("/credentials/", iamRouter.LoginWithCredentials)
			loginRoutes.POST("/telegram/", iamRouter.LoginWithTelegram)
			loginRoutes.POST("/telegram/connect/", iamRouter.ConnectTelegramUser)
		}

		userRoutes := iamRoutes.Group("/user")
		{
			meRoutes := userRoutes.Group("/me")
			{
				meRoutes.GET("/", iam.RequiresAuthentication(), iamRouter.GetUser)
				meRoutes.GET("/permissions/", iam.RequiresAuthentication(), iamRouter.ListPermissionsUser)
				meRoutes.GET("/profile/", iam.RequiresAuthentication(), iamRouter.ProfileUser)

				meRoutes.POST("/set-password/", iam.RequiresAuthentication(), iamRouter.SetPasswordUser)
				meRoutes.POST("/", iam.RequiresAuthentication(), iamRouter.UpdateUser)
			}

			userRoutes.GET("/", iam.RequiresPermission("iam.user.read"), iamRouter.ListUsers)
			userRoutes.GET("/:id/", iam.RequiresPermission("iam.user.read", "iam.user.read.:id"), iamRouter.GetUser)
			userRoutes.GET("/:id/permissions/", iam.RequiresPermission("iam.user.read", "iam.user.read.:id"), iamRouter.ListPermissionsUser)
			userRoutes.GET("/:id/profile/", iam.RequiresPermission("iam.user.read", "iam.user.read.:id"), iamRouter.ProfileUser)

			userRoutes.POST("/", iam.RequiresPermission("iam.user.create"), iamRouter.CreateUser)
			userRoutes.POST("/:id/", iam.RequiresPermission("iam.user.read", "iam.user.read.:id"), iamRouter.GetUser)
			userRoutes.POST("/:id/enabled/", iam.RequiresPermission("iam.user.write"), iamRouter.SetEnabledUser)
			userRoutes.POST("/:id/set-password/", iam.RequiresPermission("iam.user.write"), iamRouter.SetPasswordUser)

		}

		userStatusRoutes := iamRoutes.Group("/user-status")
		{
			userStatusRoutes.GET("/", iam.RequiresPermission("iam.user-status.read"), iamRouter.ListUserStatus)
			userStatusRoutes.GET("/:id/", iam.RequiresPermission("iam.user-status.read"), iamRouter.GetUserStatus)
			userStatusRoutes.GET("/:id/profile/", iam.RequiresPermission("iam.user-status.read"), iamRouter.GetUserStatusProfile)

			userStatusRoutes.POST("/", iam.RequiresPermission("iam.user-status.write"), iamRouter.CreateUserStatus)
			userStatusRoutes.POST("/:id/", iam.RequiresPermission("iam.user-status.write"), iamRouter.UpdateUserStatus)
			userStatusRoutes.POST("/:id/add-user/", iam.RequiresPermission("iam.user-status.write"), iamRouter.AddUserToUserStatus)

			userStatusRoutes.DELETE("/:id/delete/", iam.RequiresPermission("iam.user-status.write"), iamRouter.DeleteUserStatus)
		}

		userRoleRoutes := iamRoutes.Group("/user-role")
		{
			userRoleRoutes.GET("/", iam.RequiresPermission("iam.user-role.read"), iamRouter.ListUserRole)
			userRoleRoutes.GET("/:id/", iam.RequiresPermission("iam.user-role.read"), iamRouter.GetUserRole)
			userRoleRoutes.GET("/:id/profile/", iam.RequiresPermission("iam.user-role.read"), iamRouter.GetUserRoleProfile)

			userRoleRoutes.POST("/", iam.RequiresPermission("iam.user-role.write"), iamRouter.CreateUserRole)
			userRoleRoutes.POST("/:id/", iam.RequiresPermission("iam.user-role.write"), iamRouter.UpdateUserRole)
			userRoleRoutes.POST("/:id/add-user/", iam.RequiresPermission("iam.user-role.write"), iamRouter.AddUserToUserRole)
			userRoleRoutes.POST("/:id/remove-user/", iam.RequiresPermission("iam.user-role.write"), iamRouter.RemoveUserFromUserRole)

			userRoleRoutes.DELETE("/:id/delete/", iam.RequiresPermission("iam.user-role.write"), iamRouter.DeleteUserRole)
		}

		permissionRoutes := iamRoutes.Group("/permission")
		{
			permissionRoutes.GET("/", iam.RequiresPermission("iam.permission.read"), iamRouter.ListPermissions)
			permissionRoutes.GET("/:name/profile/", iam.RequiresPermission("iam.permission.read"), iamRouter.GetPermissionProfile)
		}

		permissionGroupRoutes := iamRoutes.Group("/permission-group")
		{
			permissionGroupRoutes.GET("/", iam.RequiresPermission("iam.permission-group.read"), iamRouter.ListPermissionGroups)
			permissionGroupRoutes.GET("/:id/", iam.RequiresPermission("iam.permission-group.read", "iam.permission-group.read.:id"), iamRouter.GetPermissionGroup)
			permissionGroupRoutes.GET("/:id/profile/", iam.RequiresPermission("iam.permission-group.read", "iam.permission-group.read.:id"), iamRouter.ProfilePermissionGroup)

			permissionGroupRoutes.POST("/", iam.RequiresPermission("iam.permission-group.write"), iamRouter.CreatePermissionGroup)
			permissionGroupRoutes.POST("/:id/", iam.RequiresPermission("iam.permission-group.write"), iamRouter.UpdatePermissionGroup)
			permissionGroupRoutes.POST("/:id/permission/", iam.RequiresPermission("admin"), iamRouter.EditPermissionGroupPermissions)

			permissionGroupRoutes.DELETE("/:id/", iam.RequiresPermission("iam.permission-group.write"), iamRouter.DeletePermissionGroup)
		}
	}

	sysRoutes := apiRoutes.Group("/sys")
	{
		sysRoutes.GET("/status/", iam.RequiresPermission("sys"), sysRouter.GetStatus)
		sysRoutes.GET("/installation/refresh/", sysRouter.RefreshInstallation)

		sysRoutes.POST("/logo/small/", iam.RequiresPermission("sys"), sysRouter.SetSmallLogo)
		sysRoutes.POST("/logo/large/", iam.RequiresPermission("sys"), sysRouter.SetLargeLogo)
	}
}
