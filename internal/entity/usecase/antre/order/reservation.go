package order

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/order"
)

type OrderUCInterface interface {
	order
	shop
}

type order interface {
	GetOrderByID(ctx context.Context, rsvID int64) (*rsv.TrxOrder, error)
	RegisterOrder(ctx context.Context, resrvation rsv.OrderRegistration) error
}

type shop interface {
	RegisterShop(ctx context.Context, shopRegistData rsv.ShopRegistration) error
}
