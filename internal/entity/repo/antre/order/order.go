package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/entity/order"
)

type Order interface {
	orderRegistration
	hstOrder
	dtlShop
	mapProduct
	mapShopCategory
	mapOrderProduct
}

type orderRegistration interface {
	GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error)
	InsertTrxOrder(ctx context.Context, order *order.TrxOrder) error
	UpdateTrxOrderByID(ctx context.Context, orderID int64, order *order.TrxOrder) error
	GetTrxOrderByCustomerID(ctx context.Context, userID int64, statusIDs ...int64) ([]order.TrxOrder, error)
	GetDtlOrdersByCustomerID(ctx context.Context, costomerID int64) ([]order.DtlOrder, error)
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
	GetDtlShopByShopID(ctx context.Context, shopID int64) (*order.DtlShop, error)
}

type mapProduct interface {
	InsertMapShopProduct(ctx context.Context, goodService ...order.ProductOption) error
	GetMapShopProductByShopID(ctx context.Context, shopID int64, goodServiceID ...int64) ([]order.ProductOption, error)
}

type mapShopCategory interface {
	InsertMapShopCategory(ctx context.Context, shopCategory ...order.MapShopCategory) error
}

type mapOrderProduct interface {
	InsertMapOrderProduct(ctx context.Context, orders ...*order.MapOrderProduct) error
	GetMapOrderProductByOrderID(ctx context.Context, orderID int64) ([]order.MapOrderProduct, error)
}
