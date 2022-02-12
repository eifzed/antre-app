package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
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
	if session != nil {
		session = con.DB.Slave.Context(ctx)
	}
	count, err := session.Table("ant_trx_order").Insert(&order)
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
