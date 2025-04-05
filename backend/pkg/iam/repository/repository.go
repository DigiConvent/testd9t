package iam_repository

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/pagination"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
	uuid "github.com/google/uuid"
)

type IAMRepository struct {
	db         db.DatabaseInterface
	privateKey *rsa.PrivateKey
}

type IAMRepositoryInterface interface {
	CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, core.Status)
	GetUserByEmailaddress(emailaddress string) (*iam_domain.UserRead, core.Status)
	GetUserByID(id *uuid.UUID) (*iam_domain.UserRead, core.Status)
	ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status)
	ListUsers(*iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], core.Status)
	RegisterTelegramUser(telegramId int, userId *uuid.UUID) core.Status
	SetEnabled(id *uuid.UUID, enabled bool) core.Status
	IsEnabled(id *uuid.UUID) (bool, core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) core.Status
	UserHasPermission(userId *uuid.UUID, permission string) bool
	GetUserByTelegramID(id *int) (*uuid.UUID, core.Status)
	GetTelegramID(dataString, botToken string) (*int, core.Status)
	GetUserTelegramID(id *uuid.UUID) (*int, core.Status)

	GetCredentials(emailaddress string) (*uuid.UUID, string, core.Status)
	SetCredentialPassword(id *uuid.UUID, password string) core.Status
	SetCredentialEmailaddress(id *uuid.UUID, emailaddress string) core.Status
	ResetCredentials(id *uuid.UUID) (string, core.Status)

	AddUserStatusToUser(arg *iam_domain.AddUserStatusToUser) core.Status
	CreateUserStatus(arg *iam_domain.UserStatusWrite) (*uuid.UUID, core.Status)
	DeleteUserStatus(arg *uuid.UUID) core.Status
	GetUserStatus(arg *uuid.UUID) (*iam_domain.UserStatusRead, core.Status)
	ListUserStatuses() ([]*iam_domain.UserStatusRead, core.Status)
	ListUserStatusesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasStatusRead, core.Status)
	ListUserStatusUsers(arg *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	UpdateUserStatus(id *uuid.UUID, arg *iam_domain.UserStatusWrite) core.Status

	AddUserRoleToUser(arg *iam_domain.AddUserRoleToUser) core.Status
	CreateUserRole(arg *iam_domain.UserRoleWrite) (*uuid.UUID, core.Status)
	DeleteUserRole(arg *uuid.UUID) core.Status
	GetUserRole(arg *uuid.UUID) (*iam_domain.UserRoleRead, core.Status)
	ListUserRoles() ([]*iam_domain.UserRoleRead, core.Status)
	ListUserRolesFromUser(id *uuid.UUID) ([]*iam_domain.UserHasRoleRead, core.Status)
	ListUserRoleUsers(arg *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	UpdateUserRole(id *uuid.UUID, arg *iam_domain.UserRoleWrite) core.Status

	ListUserGroups(userId *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	AddUserToPermissionGroup(permissionGroup, userId *uuid.UUID) core.Status
	CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, core.Status)
	DeletePermissionGroup(arg *uuid.UUID) core.Status
	GetPermissionGroup(arg *uuid.UUID) (*iam_domain.PermissionGroupRead, core.Status)
	ListGroupUsers(groupId *uuid.UUID) ([]*iam_domain.UserFacade, core.Status)
	ListPermissionGroupAncestors(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	ListPermissionGroupDescendants(arg *uuid.UUID) ([]*iam_domain.PermissionGroupFacade, core.Status)
	ListPermissionGroupPermissions(arg *uuid.UUID) ([]*iam_domain.PermissionFacade, core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, core.Status)
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) core.Status
	SetPermissionsForPermissionGroup(permissionGroupId *uuid.UUID, permissions []string) core.Status
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) core.Status

	ListPermissions() ([]*iam_domain.PermissionRead, core.Status)
	CreatePermission(permission *iam_domain.PermissionWrite) core.Status
	DeletePermission(name string) core.Status

	GetPrivateKey() *rsa.PrivateKey
}

func NewIamRepository(db db.DatabaseInterface) IAMRepositoryInterface {
	privateKeyPem, err := os.ReadFile(iam_setup.JwtPrivateKeyPath())
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(privateKeyPem)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return &IAMRepository{
		db:         db,
		privateKey: key,
	}
}
