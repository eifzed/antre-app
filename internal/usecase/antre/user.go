package antre

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
	authMid "github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/eifzed/antre-app/lib/utility/hash"
	"github.com/eifzed/antre-app/lib/utility/jwt"
	pkgError "github.com/pkg/errors"
)

const (
	RoleIDUser   = 6
	RoleNameUser = "User"

	MinutesInOneDay = 1440
)

func (uc *AntreUC) RegisterNewAccount(ctx context.Context, params auth.UserDetail) (*antre.RegistrationResponse, error) {
	err := params.ValidateRegistrationParam()
	if err != nil {
		return nil, err
	}
	_, err = uc.AntreDB.GetUserByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"GetUserByEmail")
	}

	// if doesnt exist then register user
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		ctx, err = uc.Transaction.Begin(ctx)
		defer uc.Transaction.Finish(ctx, &err)

		params.PasswordHashed, err = hash.HashPassword(params.Password)
		if err != nil {
			return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"HashPassword")
		}
		err = uc.AntreDB.InsertUser(ctx, &params)
		if err != nil {
			return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"InsertUser")
		}

		mapRole := &auth.MapUserRole{
			UserID: params.UserID,
			RoleID: RoleIDUser,
		}

		err = uc.AntreDB.InsertMapUserRole(ctx, mapRole)
		if err != nil {
			return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"InsertMapUserRole")
		}

		userPayload := jwt.JWTPayload{
			UserID:         params.UserID,
			Name:           params.Name,
			Email:          params.Email,
			Username:       params.Username,
			PasswordHashed: params.PasswordHashed,
			Roles:          []jwt.Role{{ID: RoleIDUser, Name: RoleNameUser}},
		}
		now := time.Now()
		var token string
		token, err = jwt.GenerateToken(userPayload, uc.Config.Secretes.Data.JWTCertificate.PrivateKey, MinutesInOneDay)
		if err != nil {
			return nil, commonerr.NewError(http.StatusInternalServerError, "JWT Token", "Failed to generate JWT Token")
		}

		response := &antre.RegistrationResponse{
			Name:       params.Name,
			Username:   params.Username,
			Email:      params.Email,
			Token:      token,
			Roles:      userPayload.Roles,
			ValidUntil: now.Add(time.Minute * MinutesInOneDay),
		}
		return response, nil
	}

	return nil, commonerr.ErrorAlreadyExist("Account", "Account with this email already exist")
}

func (uc *AntreUC) AssignNewRoleToUser(ctx context.Context, newRole string) error {

	user, _ := authMid.GetUserDetailFromContext(ctx)
	userDetail, err := uc.AntreDB.GetUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return pkgError.Wrap(err, wrapPrefixAssignNewRoleToUser+"GetUserByEmail")
	}
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		return commonerr.ErrorUnauthorized("You are not allowed to assign role")
	}

	role, err := uc.AntreDB.GetMstUserRoleByRoleName(ctx, newRole)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		pkgError.Wrap(err, wrapPrefixAssignNewRoleToUser+"getAllUserRoles")
	}

	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		return commonerr.ErrorBadRequest("Role name", "Invalid role name")
	}

	if uc.isRoleAlreadyExist(*role, user.Roles) {
		return commonerr.ErrorAlreadyExist("Role exists", "This Account already has the role")
	}
	mapRole := &auth.MapUserRole{
		UserID: userDetail.UserID,
		RoleID: role.ID,
	}
	ctx, err = uc.Transaction.Begin(ctx)
	defer uc.Transaction.Finish(ctx, &err)

	err = uc.AntreDB.InsertMapUserRole(ctx, mapRole)
	if err != nil {
		pkgError.Wrap(err, wrapPrefixAssignNewRoleToUser+"InsertMapUserRole")
	}

	return nil
}

func (uc *AntreUC) isRoleAlreadyExist(newRole antre.MstRole, userRole []jwt.Role) bool {
	for _, ur := range userRole {
		if ur.Name == newRole.Name {
			return true
		}
	}
	return false
}

func (uc *AntreUC) getAllUserRoles(ctx context.Context) ([]antre.MstRole, error) {
	// TODO: add to redis
	roles, err := uc.AntreDB.GetAllUserRoles(ctx)
	if err != nil {
		return nil, pkgError.Wrap(err, "getAllUserRoles.GetAllUserRoles")
	}
	return roles, nil
}

func (uc *AntreUC) Login(ctx context.Context, params auth.LoginParams) (*antre.RegistrationResponse, error) {
	userDetail, err := uc.AntreDB.GetUserByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"GetUserByEmail")
	}

	// if doesnt exist then register user
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		return nil, commonerr.ErrorNotFound("email")
	}
	if !hash.IsValidPasswordHash(params.Password, userDetail.PasswordHashed) {
		return nil, commonerr.ErrorUnauthorized("Invalid Password")
	}
	userRoles, err := uc.AntreDB.GetUserRolesByEmail(ctx, params.Email)
	if err != nil {
		return nil, pkgError.Wrap(err, wrapPrefixRegisterNewAccount+"GetUserRolesByEmail")
	}
	userPayload := jwt.JWTPayload{
		UserID:         userDetail.UserID,
		Name:           userDetail.Name,
		Email:          userDetail.Email,
		Username:       userDetail.Username,
		PasswordHashed: userDetail.PasswordHashed,
		Roles:          userRoles,
	}
	now := time.Now()
	var token string
	token, err = jwt.GenerateToken(userPayload, uc.Config.Secretes.Data.JWTCertificate.PrivateKey, MinutesInOneDay)
	if err != nil {
		return nil, commonerr.NewError(http.StatusInternalServerError, "JWT Token", "Failed to generate JWT Token")
	}

	response := &antre.RegistrationResponse{
		Name:       userDetail.Name,
		Username:   userDetail.Username,
		Email:      userDetail.Email,
		Token:      token,
		Roles:      userPayload.Roles,
		ValidUntil: now.Add(time.Minute * MinutesInOneDay),
	}
	return response, nil

}
