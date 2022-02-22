package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/pkg/errors"
)

func (conn *Conn) InsertMapOrderProduct(ctx context.Context, orders ...*order.MapOrderProduct) error {
	if len(orders) == 0 {
		return nil
	}
	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblMapOrderProduct).Insert(orders)
	if err != nil {
		return errors.Wrap(err, "InsertMapOrderProduct")
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) GetMapOrderProductByOrderID(ctx context.Context, orderID int64) ([]order.MapOrderProduct, error) {
	mapOrderProducts := []order.MapOrderProduct{}
	session := conn.DB.Slave.Context(ctx)
	session = session.
		Table(tblMapOrderProduct).
		Where("order_id = ?", orderID)
	count, err := session.FindAndCount(&mapOrderProducts)
	if err != nil {
		return nil, errors.Wrap(err, "GetMapOrderProductByOrderID")
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return mapOrderProducts, nil
}
