package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	"github.com/eifzed/antre-app/lib/utility/jwt"
	"github.com/go-chi/chi"
)

var (
	errInvalidToken = errors.New("Invalid authorization token")
	errUnauthorized = errors.New("User is not authorized to access this resource")
)

type AuthModule struct {
	JWTCertificate *jwt.JWTCertificate
	RouteRoles     map[string]jwt.RouteRoles
}

// fieldInfo is getter/setter value from the Info Context
type fieldInfo struct{}

type userContext struct{}

type Info struct {
	UserID int64
	Type   string
	Data   map[string]interface{}
}

func NewAuthModule(module *AuthModule) *AuthModule {
	return module
}

func (m *AuthModule) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bearerToken := r.Header.Get("Authorization")
		jwtToken, err := GetBearerToken(bearerToken)
		if err != nil {
			authHandlerError(ctx, rw, r, err)
		}
		userPayload, err := jwt.DecodeToken(jwtToken, m.JWTCertificate.PublicKey)
		if err != nil {
			authHandlerError(ctx, rw, r, err)
		}
		rCtx := chi.RouteContext(r.Context())
		if rCtx == nil {
			authHandlerError(ctx, rw, r, errors.New("Context is not Chi Context"))
		}

		route := rCtx.RouteMethod + " " + rCtx.RoutePattern()

		if !isUserAuthorized(userPayload.Roles, m.RouteRoles[route].Roles) {
			authHandlerError(ctx, rw, r, errUnauthorized)
			return
		}
		ctx = m.SetKeyValueToContext(ctx, userContext{}, auth.UserDetail{
			UserID:   userPayload.UserID,
			Name:     userPayload.Name,
			Username: userPayload.Username,
			Email:    userPayload.Email,
			Roles:    userPayload.Roles,
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func (m *AuthModule) SetKeyValueToContext(ctx context.Context, key interface{}, value interface{}) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, key, value)
}

func GetUserDetailFromContext(ctx context.Context) (auth.UserDetail, bool) {
	user, exist := ctx.Value(userContext{}).(auth.UserDetail)
	return user, exist
}

func isUserAuthorized(userRoles []jwt.Role, authorizedRoles []jwt.Role) bool {
	if len(userRoles) == 0 || len(authorizedRoles) == 0 {
		return false
	}
	for _, user := range userRoles {
		for _, auth := range authorizedRoles {
			if user.ID == auth.ID {
				return true
			}
		}
	}
	return false
}

func GetBearerToken(token string) (string, error) {
	data := strings.Split(token, "Bearer ")
	if len(data) != 2 {
		return "", errInvalidToken
	}
	return data[1], nil
}

func authHandlerError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	commonwriter.RespondError(ctx, w, err)
}
