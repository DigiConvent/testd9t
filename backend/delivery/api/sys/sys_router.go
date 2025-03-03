package sys_router

import (
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/gin-gonic/gin"
)

type SysRouterInterface interface {
	LoginWithTelegram(ctx *gin.Context)
	LoginWithCredentials(ctx *gin.Context)
}

type SysRouter struct {
	sysService  sys_service.SysServiceInterface
	postService post_service.PostServiceInterface
}

func NewSysRouter(sysService sys_service.SysServiceInterface, postService post_service.PostServiceInterface) *SysRouter {
	return &SysRouter{
		sysService:  sysService,
		postService: postService,
	}
}

func SetupSysRoutes(router *gin.Engine, sysService sys_service.SysServiceInterface, postService post_service.PostServiceInterface) {
	sysRouter := NewSysRouter(sysService, postService)

	router.POST("/status/", sysRouter.SystemStatusGet)
}
