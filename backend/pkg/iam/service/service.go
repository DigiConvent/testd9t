package iam_service

import (
	"github.com/DigiConvent/testd9t/core"
	"github.com/DigiConvent/testd9t/core/pagination"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAMServiceInterface interface {
	CreateUser(user *iam_domain.UserWrite) (*uuid.UUID, *core.Status)
	GetUser(id *uuid.UUID) (*iam_domain.UserRead, *core.Status)
	GetUserProfile(id *uuid.UUID) (*iam_domain.UserProfile, *core.Status)
	ListUsers(fs *iam_domain.UserFilterSort) (*pagination.Page[*iam_domain.UserFacade], *core.Status)
	UpdateUser(id *uuid.UUID, user *iam_domain.UserWrite) *core.Status
	SetEnabled(id *uuid.UUID, enabled bool) *core.Status

	ListUserStatuses(fs *iam_domain.UserFilterSort) ([]*iam_domain.UserStatusRead, *core.Status)
	CreateUserStatus(status *iam_domain.UserStatusWrite) (*uuid.UUID, *core.Status)
	GetUserStatus(id *uuid.UUID) (*iam_domain.UserStatusProfile, *core.Status)
	DeleteUserStatus(id *uuid.UUID) *core.Status
	UpdateUserStatus(id *uuid.UUID, status *iam_domain.UserStatusWrite) *core.Status
	AddUserStatus(status *iam_domain.AddUserStatusToUser) *core.Status

	CreatePermissionGroup(arg *iam_domain.PermissionGroupWrite) (*uuid.UUID, *core.Status)
	GetPermissionGroup(id *uuid.UUID) (*iam_domain.PermissionGroupRead, *core.Status)
	GetPermissionGroupProfile(id *uuid.UUID) (*iam_domain.PermissionGroupProfile, *core.Status)
	ListPermissionGroups() ([]*iam_domain.PermissionGroupFacade, *core.Status)
	UpdatePermissionGroup(id *uuid.UUID, arg *iam_domain.PermissionGroupWrite) *core.Status
	DeletePermissionGroup(id *uuid.UUID) *core.Status
	SetParentPermissionGroup(arg *iam_domain.PermissionGroupSetParent) *core.Status
	AddUserToPermissionGroup(permissionGroup *uuid.UUID, userId *uuid.UUID) *core.Status

	ListPermissions() ([]*iam_domain.PermissionRead, *core.Status)
	ListUserPermissions(id *uuid.UUID) ([]*iam_domain.PermissionFacade, *core.Status)
	UserHasPermission(id *uuid.UUID, permission string) bool

	// telegram and stuff
	LoginTelegramUser(body string) (*uuid.UUID, *core.Status)
	GetTelegramRegistrationCode(userId *uuid.UUID) (string, *core.Status)
	RegisterTelegramUser(telegramId int, email, code string) *core.Status

	// password and stuff
	ResetPassword(email string) (string, *core.Status)
	SetUserPassword(id *uuid.UUID, password string) *core.Status
	LoginUser(email, password string) (*uuid.UUID, *core.Status)

	// jwt stuff
	GenerateJwt(userId *uuid.UUID) (string, *core.Status)
	VerifyJwt(token string) (*uuid.UUID, *core.Status)
}

type IAMService struct {
	IAMRepository iam_repository.IAMRepositoryInterface
}

func NewIAMService(userRepository iam_repository.IAMRepositoryInterface) IAMServiceInterface {
	return &IAMService{
		IAMRepository: userRepository,
	}
}

func hashedPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}
