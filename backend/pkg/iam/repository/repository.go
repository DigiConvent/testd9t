package iam_repository

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	uuid "github.com/google/uuid"
)

type IAMRepository struct {
	DB db.DatabaseInterface
}

type IAMRepositoryInterface interface {
	VerifyTelegramUser(body string) (*uuid.UUID, core.Status)
	GetTelegramRegistrationCode(userId *uuid.UUID) (string, core.Status)
	RegisterTelegramUser(telegramId int, email, code string) core.Status
	CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, core.Status)
	GetUserByID(id *uuid.UUID) (*iam_domain.UserRead, core.Status)
	ListUsers() ([]*iam_domain.UserFacade, core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status
	ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status)
	SetEnabled(id *uuid.UUID, enabled bool) core.Status

	CreateUserStatus(arg *iam_domain.UserStatusWrite) (*uuid.UUID, core.Status)
	ListUserStatuses() ([]*iam_domain.UserStatusRead, core.Status)
	GetUserStatus(arg *uuid.UUID) (*iam_domain.UserStatusRead, core.Status)
	ListUserStatusUsers(arg *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	UpdateUserStatus(id *uuid.UUID, arg *iam_domain.UserStatusWrite) core.Status
	DeleteUserStatus(arg *uuid.UUID) core.Status
	AddUserStatusToUser(arg *iam_domain.AddUserStatusToUser) core.Status
	ListUserStatusesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasStatusRead, core.Status)

	CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, core.Status)
	GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroupProfile, core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupRead, core.Status)
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) core.Status
	DeletePermissionGroup(arg *uuid.UUID) core.Status
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) core.Status
	SetPermissionsForPermissionGroup(permissionGroupId *uuid.UUID, permissions []*uuid.UUID) core.Status
	ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	ListUserGroups(userId *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)

	ListPermissions() ([]*iam_domain.PermissionRead, core.Status)
	UserHasPermission(userId *uuid.UUID, permission string) (bool, core.Status)
}

func NewIAMRepository(db db.DatabaseInterface) IAMRepositoryInterface {
	return &IAMRepository{
		DB: db,
	}
}
