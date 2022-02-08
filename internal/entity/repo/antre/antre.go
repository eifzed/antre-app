package antre

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/lib/utility/jwt"
)

type AntreDBInterface interface {
	user
	mapUserRole
	mstUser
}

type user interface {
	GetUserByEmail(ctx context.Context, email string) (*auth.UserDetail, error)
	GetUserByUserID(ctx context.Context, userID int64) (*auth.UserDetail, error)
	InsertUser(ctx context.Context, userParam *auth.UserDetail) error
	UpdateUserByUserID(ctx context.Context, userID int64, userParam *auth.UserDetail) error
	DeleteUserByUserID(ctx context.Context, userID int64) error
}

type mapUserRole interface {
	GetMapUserRoleByEmail(ctx context.Context, email string) ([]auth.MapUserRole, error)
	GetUserRolesByEmail(ctx context.Context, email string) ([]jwt.Role, error)
	InsertMapUserRole(ctx context.Context, user *auth.MapUserRole) error
}

type mstUser interface {
	GetAllUserRoles(ctx context.Context) ([]MstRole, error)
	GetMstUserRoleByRoleName(ctx context.Context, roleName string) (*MstRole, error)
}
