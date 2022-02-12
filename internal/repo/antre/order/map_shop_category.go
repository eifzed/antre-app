package order

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (conn *Conn) InsertMapShopCategory(ctx context.Context, shopCategory ...rsv.MapShopCategory) error {
	if len(shopCategory) == 0 {
		return nil
	}

	session := getSession(ctx)
	if session == nil {
		session = conn.DB.Master.Context(ctx)
	}
	count, err := session.Table(tblMapShopCategory).Insert(shopCategory)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}
