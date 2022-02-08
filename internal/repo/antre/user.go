package antre

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
	"github.com/eifzed/antre-app/lib/utility/jwt"
	"github.com/pkg/errors"
)

type Conn struct {
	DB *db.Connection
	// Gocrypt *gocrypt.Option
}

func NewDBConnection(conn *db.Connection) *Conn {
	return &Conn{
		DB: conn,
	}
}

var (
	getSession = db.GetDBSession
)

const (
	wrapPrefix                      = "repo.antre."
	wrapPrefixGetUserByEmail        = wrapPrefix + "GetUserByEmail."
	wrapPrefixGetUserByUserID       = wrapPrefix + "GetUserByUserID."
	wrapPrefixInsertUser            = wrapPrefix + "InsertUser."
	wrapPrefixUpdateUserByUserID    = wrapPrefix + "UpdateUserByUserID."
	wrapPrefixDeleteUserByUserID    = wrapPrefix + "DeleteUserByUserID."
	wrapPrefixGetMapUserRoleByEmail = wrapPrefix + "GetMapUserRoleByEmail."
	wrapPrefixInsertMapUserRole     = wrapPrefix + "InsertMapUserRole."
	wrapPrefixLoginUser             = wrapPrefix + "LoginUser."
)

func (conn *Conn) GetUserByEmail(ctx context.Context, email string) (*auth.UserDetail, error) {
	session := conn.DB.Slave.Context(ctx).Table("ant_mst_user")
	user := auth.UserDetail{}
	has, err := session.Where("email ILIKE ?", email).Get(&user)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetUserByEmail)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return &user, nil
}

func (conn *Conn) GetUserByUserID(ctx context.Context, userID int64) (*auth.UserDetail, error) {
	session := conn.DB.Slave.Context(ctx).Table("ant_mst_user")
	user := auth.UserDetail{}
	has, err := session.Where("user_id = ?", userID).Get(&user)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetUserByUserID)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return &user, nil
}

func (conn *Conn) InsertUser(ctx context.Context, userParam *auth.UserDetail) error {
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}

	count, err := session.Table("ant_mst_user").Insert(userParam)
	if err != nil {
		return errors.Wrap(err, wrapPrefixInsertUser)
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil

}

func (conn *Conn) UpdateUserByUserID(ctx context.Context, userID int64, userParam *auth.UserDetail) error {
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table("ant_mst_user").Where("user_id = ?", userID).Update(userParam)
	if err != nil {
		return errors.Wrap(err, wrapPrefixUpdateUserByUserID)
	}
	if count == 0 {
		return databaseerr.ErrorNoUpdate
	}
	return nil
}

func (conn *Conn) DeleteUserByUserID(ctx context.Context, userID int64) error {
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table("ant_mst_user").
		Where("user_id = ?", userID).
		Delete(&auth.UserDetail{})
	if err != nil {
		return errors.Wrap(err, wrapPrefixDeleteUserByUserID)
	}
	if count == 0 {
		return databaseerr.ErrorNoDelete
	}
	return nil
}

func (conn *Conn) GetMapUserRoleByEmail(ctx context.Context, email string) ([]auth.MapUserRole, error) {
	mapUserRole := []auth.MapUserRole{}
	session := conn.DB.Slave.Context(ctx).Table("ant_map_user_role")
	has, err := session.Where("email = ?", email).Get(&mapUserRole)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapUserRoleByEmail)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return mapUserRole, nil
}

func (conn *Conn) GetUserRolesByEmail(ctx context.Context, email string) ([]jwt.Role, error) {
	userRole := []jwt.Role{}
	session := conn.DB.Slave.Context(ctx).
		Select("amr.role_id, amr.role_name").
		Table("ant_mst_user").Alias("amu").
		Join("LEFT", "ant_map_user_role amur", "amur.user_id = amu.user_id").
		Join("LEFT", "ant_mst_role amr", "amr.role_id = amur.role_id")
	count, err := session.Where("amu.email ILIKE ?", email).FindAndCount(&userRole)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapUserRoleByEmail)
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return userRole, nil
}

func (conn *Conn) InsertMapUserRole(ctx context.Context, user *auth.MapUserRole) error {
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table("ant_map_user_role").Insert(user)
	if err != nil {
		return errors.Wrap(err, wrapPrefixInsertMapUserRole)
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) GetAllUserRoles(ctx context.Context) ([]antre.MstRole, error) {
	roles := []antre.MstRole{}
	session := conn.DB.Slave.Context(ctx).Table("ant_mst_role")
	has, err := session.Get(&roles)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapUserRoleByEmail)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return roles, nil
}

func (conn *Conn) GetMstUserRoleByRoleName(ctx context.Context, roleName string) (*antre.MstRole, error) {
	role := antre.MstRole{}
	session := conn.DB.Slave.Context(ctx).Table("ant_mst_role").Where("role_name = ?", roleName)
	has, err := session.Get(&role)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapUserRoleByEmail)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return &role, nil
}
