package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/order"
)

type OrderUCInterface interface {
	orders
	shop
	product
}

type orders interface {
	GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error)
	RegisterOrder(ctx context.Context, resrvation order.OrderRegistration) error
	GetCustomerOrders(ctx context.Context) (order.DtlOrderList, error)
}

type shop interface {
	RegisterShop(ctx context.Context, shopRegistData order.ShopRegistration) error
}

type product interface {
	GetProductsListByShopID(ctx context.Context, shopID int64) (*order.ProductsList, error)
}
