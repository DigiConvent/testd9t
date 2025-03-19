package sys_router

import (
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/gin-gonic/gin"
)

type SysRouterInterface interface {
	SetSmallLogo(c *gin.Context)
	SetLargeLogo(c *gin.Context)

	GetStatus(c *gin.Context)
	RefreshInstallation(c *gin.Context)
}

type SysRouter struct {
	sysService sys_service.SysServiceInterface
}

func NewSysRouter(sysService sys_service.SysServiceInterface) SysRouterInterface {
	return &SysRouter{
		sysService: sysService,
	}
}

func SetupSysRoutes(router *gin.Engine, sysService sys_service.SysServiceInterface) {
	sysRouter := NewSysRouter(sysService)

	router.POST("/status/", sysRouter.GetStatus)
}
