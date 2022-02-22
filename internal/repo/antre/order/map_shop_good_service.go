package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/pkg/errors"
)

func (conn *Conn) InsertMapShopProduct(ctx context.Context, goodService ...order.ProductOption) error {
	if len(goodService) == 0 {
		return nil
	}

	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblMapShopProduct).Insert(goodService)
	if err != nil {
		return errors.Wrap(err, "InsertMapShopProduct")
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) GetMapShopProductByShopID(ctx context.Context, shopID int64, goodServiceID ...int64) ([]order.ProductOption, error) {
	orders := []order.ProductOption{}
	session := conn.DB.Slave.Context(ctx)
	session = session.
		Table(tblMapShopProduct).
		Where("shop_id = ?", shopID)
	if len(goodServiceID) > 0 {
		session.In("id", goodServiceID)
	}
	count, err := session.FindAndCount(&orders)
	if err != nil {
		return nil, errors.Wrap(err, "GetMapShopProductByShopID")
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return orders, nil
}
