package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/order"
	"github.com/pkg/errors"
)

func (uc *OrderUC) GetProductsListByShopID(ctx context.Context, shopID int64) (*order.ProductsList, error) {

	result, err := uc.OrderDB.GetProductsListByShopID(ctx, shopID)
	if err != nil {
		return nil, errors.Wrap(err, wrapPrefixGetProductsListByShopID)
	}
	return result, nil
}
