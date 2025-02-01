package sys_router

import sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"

type SysRouter struct {
	SysService sys_service.SysServiceInterface
}

func NewSysRouter(sysService sys_service.SysServiceInterface) *SysRouter {
	return &SysRouter{
		SysService: sysService,
	}
}
