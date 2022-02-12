package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
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
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}