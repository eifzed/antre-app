package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/pkg/errors"
)

func (conn *Conn) InsertMapShopGoodService(ctx context.Context, goodService ...order.GoodServiceOption) error {
	if len(goodService) == 0 {
		return nil
	}

	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblMapShopGoodService).Insert(goodService)
	if err != nil {
		return errors.Wrap(err, wrapPrefixInsertMapShopGoodService)
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) GetMapShopGoodServiceByShopID(ctx context.Context, shopID int64, goodServiceID ...int64) ([]order.GoodServiceOption, error) {
	orders := []order.GoodServiceOption{}
	session := conn.DB.Slave.Context(ctx)
	session = session.
		Table(tblMapShopGoodService).
		Where("shop_id = ?", shopID)
	if len(goodServiceID) > 0 {
		session.In("id", goodServiceID)
	}
	count, err := session.FindAndCount(&orders)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetMapShopGoodServiceByShopID)
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return orders, nil
}
