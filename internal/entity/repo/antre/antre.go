package antre

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/auth"
)

type AntreDBInterface interface {
	user
}

type user interface {
	GetUserByEmail(ctx context.Context, email string) (*auth.User, error)
	GetUserByUserID(ctx context.Context, userID int64) (*auth.User, error)
	InsertUser(ctx context.Context, userParam *auth.User) error
	UpdateUserByUserID(ctx context.Context, userID int64, userParam *auth.User) error
	DeleteUserByUserID(ctx context.Context, userID int64) error
}
