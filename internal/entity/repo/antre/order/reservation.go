package order

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/order"
)

type Order interface {
	order
	hstOrder
	dtlShop
	mapGoodService
	mapShopCategory
}

type order interface {
	GetOrderByID(ctx context.Context, rsvID int64) (*rsv.TrxOrder, error)
	InsertTrxOrder(ctx context.Context, order *rsv.TrxOrder) error
	UpdateTrxOrderByID(ctx context.Context, rsvID int64, order *rsv.TrxOrder) error
}

type hstOrder interface {
	GetHstOrderByRsvID(ctx context.Context, rsvID int64) (*rsv.HstOrder, error)
	GetHstOrderByHstID(ctx context.Context, hstOrderID int64) (*rsv.HstOrder, error)
	InsertHstOrder(ctx context.Context, hstOrder *rsv.HstOrder) error
	UpdateHstOrderByHstID(ctx context.Context, hstOrderID int64, hstOrder *rsv.HstOrder) error
}

type dtlShop interface {
	GetDtlShopByOwnerID(ctx context.Context, ownerID int64) (*rsv.DtlShop, error)
	InsertDtlShopByOwnerID(ctx context.Context, shopData *rsv.DtlShop) error
}

type mapGoodService interface {
	InsertMapShopGoodService(ctx context.Context, goodService ...rsv.GoodServiceOption) error
}

type mapShopCategory interface {
	InsertMapShopCategory(ctx context.Context, shopCategory ...rsv.MapShopCategory) error
}
