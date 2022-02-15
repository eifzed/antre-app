package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
	"github.com/pkg/errors"
)

var (
	getSession = db.GetDBSession
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

func (con *Conn) GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error) {
	data := &order.TrxOrder{}
	session := con.DB.Slave.Context(ctx).Table("ant_trx_order")
	_, err := session.Where("order_id = ?", orderID).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (con *Conn) InsertTrxOrder(ctx context.Context, order *order.TrxOrder) error {
	session := getSession(ctx)
	if session == nil {
		session = con.DB.Slave.Context(ctx)
	}
	count, err := session.Table("ant_trx_order").Insert(order)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (con *Conn) UpdateTrxOrderByID(ctx context.Context, orderID int64, order *order.TrxOrder) error {
	session := getSession(ctx)
	if session != nil {
		session = con.DB.Slave.Context(ctx)
	}
	count, err := session.Table("ant_trx_order").
		Where("order_id = ?", orderID).
		Update(&order)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (con *Conn) GetTrxOrderByCustomerID(ctx context.Context, userID int64, statusIDs ...int64) ([]order.TrxOrder, error) {
	data := []order.TrxOrder{}
	session := con.DB.Slave.Context(ctx).
		Table(tblTrxOrder).Alias("trx").
		Join("LEFT", tblMapOrderGoodService+" map", "trx.order_id = map.order_id")
	session = session.Where("customer_id = ?", userID)
	if len(statusIDs) > 0 {
		session.In("status_id", statusIDs)
	}

	count, err := session.FindAndCount(data)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return data, nil
}

func (conn *Conn) GetDtlOrdersByCustomerID(ctx context.Context, costomerID int64) ([]order.DtlOrder, error) {
	result := []order.DtlOrder{}
	session := conn.DB.Slave.Context(ctx).
		Table(tblTrxOrder).Alias("trx").
		Join("LEFT", tblDtlShop+" shop", "shop.shop_id = trx.shop_id").
		Join("LEFT", tblMstStatus+" status", "status.status_id = trx.status_id").
		Select("trx.shop_id, shop.shop_name, status.status_name, trx.customer_note, trx.shop_note, trx.create_time, trx.order_id")
	count, err := session.Where("customer_id = ?", costomerID).FindAndCount(&result)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetOrdersByCustomerID)
	}
	if count == 0 {
		return nil, databaseerr.ErrorDataNotFound
	}
	return result, nil
}
