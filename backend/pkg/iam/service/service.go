package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	"github.com/google/uuid"
)

type IAMServiceInterface interface {
	// user relevant methods
	CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status)
	ReadUser(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status)
	ListUsers() ([]*iam_domain.UserFacade, *core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) *core.Status
	LoginTelegramUser(body string) (*uuid.UUID, *core.Status)
	GetTelegramRegistrationCode(userId *uuid.UUID) (string, *core.Status)
	RegisterTelegramUser(telegramId int, email, code string) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status

	// user status relevant methods
	ListUserStatuses() ([]*iam_domain.UserStatusRead, *core.Status)
	CreateUserStatus(status *iam_domain.UserStatusWrite) (*uuid.UUID, *core.Status)
	GetUserStatus(id *uuid.UUID) (*iam_domain.UserStatusProfile, *core.Status)
	DeleteUserStatus(id *uuid.UUID) *core.Status
	UpdateUserStatus(id *uuid.UUID, status *iam_domain.UserStatusWrite) *core.Status
	AddUserStatus(status *iam_domain.AddUserStatusToUser) *core.Status

	// permission group relevant methods
	CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, *core.Status)
	GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupRead, *core.Status)
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) *core.Status
	DeletePermissionGroup(id *uuid.UUID) *core.Status
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status

	// permission relevant methods
	ListPermissions() ([]*iam_domain.PermissionRead, *core.Status)
	ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, *core.Status)
	UserHasPermission(id *uuid.UUID, permission string) (bool, *core.Status)
}

type IAMService struct {
	IAMRepository iam_repository.IAMRepositoryInterface
}

func NewIAMService(userRepository iam_repository.IAMRepositoryInterface) IAMServiceInterface {
	return &IAMService{
		IAMRepository: userRepository,
	}
}
