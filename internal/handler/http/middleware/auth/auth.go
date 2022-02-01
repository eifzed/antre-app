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

func (m *AuthModule) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		bearerToken := r.Header.Get("Authorization")
		jwtToken, err := GetBearerToken(bearerToken)
		if err != nil {
			authHandlerError(ctx, rw, r, err)
		}
		fmt.Println(jwtToken)

	})
}

func validateJWTToken(token string, publicKey string) ([]byte, error) {
	tokenArray := strings.Split(token, ".")
	if len(tokenArray) != 3 {
		return nil, errInvalidToken
	}
	return nil, nil
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
