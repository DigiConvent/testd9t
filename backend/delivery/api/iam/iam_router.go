package iam_router

import (
	router_utils "github.com/DigiConvent/testd9t/delivery/api/util"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/gin-gonic/gin"
)

type IamRouterInterface interface {
	LoginWithTelegram(ctx *gin.Context)
	LoginWithCredentials(ctx *gin.Context)
	LoginWithJwt(ctx *gin.Context)

	CreatePermissionGroup(ctx *gin.Context)
	DeletePermissionGroup(ctx *gin.Context)
	GetPermissionGroup(ctx *gin.Context)
	ListPermissionGroups(ctx *gin.Context)
	ProfilePermissionGroup(ctx *gin.Context)
	SetParentPermissionGroup(ctx *gin.Context)
	UpdatePermissionGroup(ctx *gin.Context)
	EditPermissionGroupPermissions(ctx *gin.Context)

	ListPermissions(ctx *gin.Context)
	GetPermissionProfile(ctx *gin.Context)

	CreateUserStatus(ctx *gin.Context)
	GetUserStatus(ctx *gin.Context)
	GetUserStatusProfile(ctx *gin.Context)
	ListUserStatus(ctx *gin.Context)
	UpdateUserStatus(ctx *gin.Context)
	AddUserToUserStatus(ctx *gin.Context)
	DeleteUserStatus(ctx *gin.Context)

	CreateUserRole(ctx *gin.Context)
	GetUserRole(ctx *gin.Context)
	GetUserRoleProfile(ctx *gin.Context)
	ListUserRole(ctx *gin.Context)
	UpdateUserRole(ctx *gin.Context)
	AddUserToUserRole(ctx *gin.Context)
	RemoveUserFromUserRole(ctx *gin.Context)
	DeleteUserRole(ctx *gin.Context)

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

	ConnectTelegramUser(ctx *gin.Context)
}

type IamRouter struct {
	iamService iam_service.IAMServiceInterface
	sysService sys_service.SysServiceInterface
}

// AddUserToUserStatus implements IamRouterInterface.
func (router *IamRouter) AddUserToUserStatus(ctx *gin.Context) {
	panic("unimplemented")
}

// DeleteUserStatus implements IamRouterInterface.
func (router *IamRouter) DeleteUserStatus(ctx *gin.Context) {
	panic("unimplemented")
}

// GetUserRole implements IamRouterInterface.
func (router *IamRouter) GetUserRole(ctx *gin.Context) {
	id := router_utils.GetId(ctx)

	userRole, status := router.iamService.GetUserRole(id)

	if status != nil && status.Err() {
		ctx.JSON(status.Code, gin.H{
			"error": status.Message,
		})
	} else {
		ctx.JSON(status.Code, userRole)
	}
}

// HasPermissionUser implements IamRouterInterface.
func (router *IamRouter) HasPermissionUser(ctx *gin.Context) {
	panic("unimplemented")
}

// ListUserRole implements IamRouterInterface.
func (router *IamRouter) ListUserRole(ctx *gin.Context) {
	panic("unimplemented")
}

// ResetPasswordUser implements IamRouterInterface.
func (router *IamRouter) ResetPasswordUser(ctx *gin.Context) {
	panic("unimplemented")
}

// SetParentPermissionGroup implements IamRouterInterface.
func (router *IamRouter) SetParentPermissionGroup(ctx *gin.Context) {
	panic("unimplemented")
}

// UpdateUserRole implements IamRouterInterface.
func (router *IamRouter) UpdateUserRole(ctx *gin.Context) {
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
