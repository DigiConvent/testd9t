package sys_router

import (
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

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
