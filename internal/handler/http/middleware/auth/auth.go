package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/eifzed/antre-app/lib/common/commonwriter"
	"github.com/eifzed/antre-app/lib/utility/jwt"
)

var (
	errInvalidToken = errors.New("Invalid authorization token")
	errUnauthorized = errors.New("User is not authorized to access this resource")
)

type AuthModule struct {
	JWTCertificate *jwt.JWTCertificate
	RouteRoles     map[string]*jwt.RouteRoles
}

func NewAuthModule(module *AuthModule) *AuthModule {
	return module
}

var (
	roles = []jwt.Role{{ID: 123, Name: "dev"}}
)

func (m *AuthModule) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bearerToken := r.Header.Get("Authorization")
		jwtToken, err := GetBearerToken(bearerToken)
		if err != nil {
			authHandlerError(ctx, rw, r, err)
		}
		fmt.Println(jwtToken)
		userPayload, err := jwt.DecodeToken(jwtToken, m.JWTCertificate.PublicKey)
		if err != nil {
			authHandlerError(ctx, rw, r, err)
		}

		if !isUserAuthorized(userPayload.Roles, roles) {
			authHandlerError(ctx, rw, r, errUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
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
	commonwriter.RespondError(ctx, w, http.StatusUnauthorized, err.Error())
}
