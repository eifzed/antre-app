package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/pkg/errors"
)

func (conn *Conn) InsertMapOrderGoodService(ctx context.Context, orders ...*order.MapOrderGoodService) error {
	if len(orders) == 0 {
		return nil
	}
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblMapOrderGoodService).Insert(orders)
	if err != nil {
		return errors.Wrap(err, wrapPrefixInsertMapOrderGoodService)
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) GetMapOrderGoodServiceByOrderID(ctx context.Context, orderID int64) ([]order.MapOrderGoodService, error) {
	mapOrderGoodServices := []order.MapOrderGoodService{}
	session := conn.DB.Slave.Context(ctx)
	session = session.
		Table(tblMapOrderGoodService).
		Where("order_id = ?", orderID)
	count, err := session.FindAndCount(&mapOrderGoodServices)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapShopGoodServiceByShopID)
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return mapOrderGoodServices, nil
}
