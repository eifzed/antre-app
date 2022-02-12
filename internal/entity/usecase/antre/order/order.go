package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/order"
)

type OrderUCInterface interface {
	orders
	shop
}

type orders interface {
	GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error)
	RegisterOrder(ctx context.Context, resrvation order.OrderRegistration) error
}

type shop interface {
	RegisterShop(ctx context.Context, shopRegistData order.ShopRegistration) error
}
