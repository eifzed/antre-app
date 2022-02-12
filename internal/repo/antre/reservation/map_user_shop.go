package reservation

import (
	"context"

	"github.com/pkg/errors"

	"github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (conn *Conn) GetDtlShopByOwnerID(ctx context.Context, ownerID int64) (*reservation.DtlShop, error) {
	dtlShop := &reservation.DtlShop{}
	session := conn.DB.Slave.Context(ctx).Table(tblDtlShop)
	has, err := session.Where("owner_id = ?", ownerID).Get(dtlShop)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetDtlShopByOwnerID)
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return dtlShop, nil
}

func (conn *Conn) InsertDtlShopByOwnerID(ctx context.Context, shopData *reservation.DtlShop) error {
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Slave.Context(ctx)
	}
	count, err := session.Table(tblDtlShop).InsertOne(shopData)
	if err != nil {
		return errors.Wrap(err, wrapPrefixGetDtlShopByOwnerID)
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}
