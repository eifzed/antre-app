package order

import (
	"context"

	order "github.com/eifzed/antre-app/internal/entity/order"
	"github.com/pkg/errors"
)

func (conn *Conn) GetProductsListByShopID(ctx context.Context, shopID int64) (*order.ProductsList, error) {
	products := []order.Product{}
	session := conn.DB.Slave.Context(ctx)
	count, err := session.
		Select("").
		Table(tblMapShopProduct).
		Where("shop_id = ?", shopID).
		FindAndCount(&products)
	if err != nil {
		return nil, errors.Wrap(err, "GetProductsListByShopID")
	}
	result := &order.ProductsList{
		Total:    count,
		Products: products,
	}
	return result, nil
}
