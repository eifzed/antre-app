package antre

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
)

type AntreUCInterface interface {
	user
}

type user interface {
	RegisterNewAccount(ctx context.Context, userInfo auth.UserDetail) (*antre.RegistrationResponse, error)
	AssignNewRoleToUser(ctx context.Context, newRole string) error
	Login(ctx context.Context, params auth.LoginParams) (*antre.RegistrationResponse, error)
}
