package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (conn *Conn) GetUserByEmail(ctx context.Context, email string) (*rsv.User, error) {
	session := conn.DB.Slave.Table("ant_mst_user")
	user := rsv.User{}
	has, err := session.Where("email = ?", email).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return &user, nil
}

func (conn *Conn) GetUserByUserID(ctx context.Context, userID int64) (*rsv.User, error) {
	session := conn.DB.Slave.Context(ctx).Table("ant_mst_user")
	user := rsv.User{}
	has, err := session.Where("user_id = ?", userID).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return &user, nil
}

func (conn *Conn) InsertUser(ctx context.Context, userParam *rsv.User) error {
	session := conn.DB.Master.Context(ctx).Table("ant_mst_user")
	count, err := session.Insert(userParam)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil

}

func (conn *Conn) UpdateUserByUserID(ctx context.Context, userID int64, userParam *rsv.User) error {
	session := conn.DB.Master.Context(ctx).Table("ant_mst_user")
	count, err := session.Where("user_id = ?", userID).Update(userParam)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoUpdate
	}
	return nil
}

func (conn *Conn) DeleteUserByUserID(ctx context.Context, userID int64) error {
	session := conn.DB.Master.Context(ctx).Table("ant_mst_user")
	count, err := session.Where("user_id = ?", userID).Delete(&rsv.User{})
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoDelete
	}
	return nil
}