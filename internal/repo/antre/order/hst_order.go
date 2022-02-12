package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (conn *Conn) GetHstOrderByHstID(ctx context.Context, hstOrderID int64) (*order.HstOrder, error) {
	hstOrder := &order.HstOrder{}
	session := conn.DB.Master.Context(ctx).Table(tblHstOrder)
	has, err := session.Where("hst_order_id = ?", hstOrderID).Get(hstOrder)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return hstOrder, nil
}

func (conn *Conn) GetHstOrderByOrderID(ctx context.Context, orderID int64) (*order.HstOrder, error) {
	hstOrder := &order.HstOrder{}
	session := conn.DB.Master.Context(ctx).Table(tblHstOrder)
	has, err := session.Where("order_id = ?", orderID).Get(hstOrder)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return hstOrder, nil
}

func (conn *Conn) InsertHstOrder(ctx context.Context, hstOrder *order.HstOrder) error {
	session := getSession(ctx)
	if session != nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblHstOrder).Insert(hstOrder)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) UpdateHstOrderByHstID(ctx context.Context, hstOrderID int64, hstOrder *order.HstOrder) error {
	session := conn.DB.Master.Context(ctx).Table(tblHstOrder)
	count, err := session.Where("hst_order_id = ?", hstOrderID).Update(hstOrder)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoUpdate
	}
	return nil
}
