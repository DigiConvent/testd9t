package iam_router

import (
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/gin-gonic/gin"
)

type IamRouterInterface interface {
	LoginWithTelegram(ctx *gin.Context)
	LoginWithCredentials(ctx *gin.Context)

	AddUserPermissionGroup(ctx *gin.Context)
	CreatePermissionGroup(ctx *gin.Context)
	DeletePermissionGroup(ctx *gin.Context)
	GetPermissionGroup(ctx *gin.Context)
	ListPermissionGroups(ctx *gin.Context)
	ProfilePermissionGroup(ctx *gin.Context)
	SetParentPermissionGroup(ctx *gin.Context)
	UpdatePermissionGroup(ctx *gin.Context)

	ListPermissions(ctx *gin.Context)

	CreateUserStatus(ctx *gin.Context)
	DeleteUserStatus(ctx *gin.Context)
	GetUserStatus(ctx *gin.Context)
	ListUserStatus(ctx *gin.Context)
	UpdateUserStatus(ctx *gin.Context)
	AddUserStatusUser(ctx *gin.Context)

	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	HasPermissionUser(ctx *gin.Context)
	ListPermissionsUser(ctx *gin.Context)
	ListUsers(ctx *gin.Context)
	ProfileUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)

	ResetPasswordUser(ctx *gin.Context)
	SetEnabledUser(ctx *gin.Context)
	SetPasswordUser(ctx *gin.Context)

	GetCodeForTelegramUser(ctx *gin.Context)
	ConnectTelegramUser(ctx *gin.Context)
}

type IamRouter struct {
	iamService iam_service.IAMServiceInterface
	sysService sys_service.SysServiceInterface
}

// AddUserPermissionGroup implements IamRouterInterface.
func (router *IamRouter) AddUserPermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// AddUserStatusUser implements IamRouterInterface.
func (router *IamRouter) AddUserStatusUser(ctx *gin.Context) {
	panic("unimplemented")
}

// CreatePermissionGroup implements IamRouterInterface.
func (router *IamRouter) CreatePermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// DeletePermissionGroup implements IamRouterInterface.
func (router *IamRouter) DeletePermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteUserStatus implements IamRouterInterface.
func (router *IamRouter) DeleteUserStatus(ctx *gin.Context) {
	panic("unimplemented")
}

// GetCodeForTelegramUser implements IamRouterInterface.
func (router *IamRouter) GetCodeForTelegramUser(ctx *gin.Context) {
	panic("unimplemented")
}

// GetUser implements IamRouterInterface.
func (router *IamRouter) GetUser(ctx *gin.Context) {
	panic("unimplemented")
}

// GetUserStatus implements IamRouterInterface.
func (router *IamRouter) GetUserStatus(ctx *gin.Context) {
	panic("unimplemented")
}

// HasPermissionUser implements IamRouterInterface.
func (router *IamRouter) HasPermissionUser(ctx *gin.Context) {
	panic("unimplemented")
}

// ProfileUser implements IamRouterInterface.
func (router *IamRouter) ProfileUser(ctx *gin.Context) {
	panic("unimplemented")
}

// ResetPasswordUser implements IamRouterInterface.
func (router *IamRouter) ResetPasswordUser(ctx *gin.Context) {
	panic("unimplemented")
}

// SetEnabledUser implements IamRouterInterface.
func (router *IamRouter) SetEnabledUser(ctx *gin.Context) {
	panic("unimplemented")
}

// SetParentPermissionGroup implements IamRouterInterface.
func (router *IamRouter) SetParentPermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// SetPasswordUser implements IamRouterInterface.
func (router *IamRouter) SetPasswordUser(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdatePermissionGroup implements IamRouterInterface.
func (router *IamRouter) UpdatePermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateUserStatus implements IamRouterInterface.
func (router *IamRouter) UpdateUserStatus(ctx *gin.Context) {
	panic("unimplemented")
}

func NewIamRouter(iamService iam_service.IAMServiceInterface, sysService sys_service.SysServiceInterface) IamRouterInterface {
	return &IamRouter{
		iamService: iamService,
		sysService: sysService,
	}
}
