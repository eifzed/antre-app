package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
)

type ReservationUCInterface interface {
	reservation
	shop
}

type reservation interface {
	GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error)
	RegisterReservation(ctx context.Context, resrvation *rsv.TrxReservation) error
}

type shop interface {
	RegisterShop(ctx context.Context, shopRegistData rsv.ShopRegistration) error
}
