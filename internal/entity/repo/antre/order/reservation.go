package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/order"
)

type Order interface {
	orderRegistration
	hstOrder
	dtlShop
	mapGoodService
	mapShopCategory
}

type orderRegistration interface {
	GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error)
	InsertTrxOrder(ctx context.Context, order *order.TrxOrder) error
	UpdateTrxOrderByID(ctx context.Context, orderID int64, order *order.TrxOrder) error
}

type hstOrder interface {
	GetHstOrderByOrderID(ctx context.Context, orderID int64) (*order.HstOrder, error)
	GetHstOrderByHstID(ctx context.Context, hstOrderID int64) (*order.HstOrder, error)
	InsertHstOrder(ctx context.Context, hstOrder *order.HstOrder) error
	UpdateHstOrderByHstID(ctx context.Context, hstOrderID int64, hstOrder *order.HstOrder) error
}

type dtlShop interface {
	GetDtlShopByOwnerID(ctx context.Context, ownerID int64) (*order.DtlShop, error)
	InsertDtlShopByOwnerID(ctx context.Context, shopData *order.DtlShop) error
}

type mapGoodService interface {
	InsertMapShopGoodService(ctx context.Context, goodService ...order.GoodServiceOption) error
}

type mapShopCategory interface {
	InsertMapShopCategory(ctx context.Context, shopCategory ...order.MapShopCategory) error
}
