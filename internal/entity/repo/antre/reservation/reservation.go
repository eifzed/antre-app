package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
)

type Reservation interface {
	reservation
	hstReservation
}

type reservation interface {
	GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error)
	InsertTrxReservation(ctx context.Context, reservation *rsv.TrxReservation) error
	UpdateTrxReservationByID(ctx context.Context, rsvID int64, reservation *rsv.TrxReservation) error
}

type hstReservation interface {
	GetHstReservationByRsvID(ctx context.Context, rsvID int64) (*rsv.HstReservation, error)
	GetHstReservationByHstID(ctx context.Context, hstReservationID int64) (*rsv.HstReservation, error)
	InsertHstReservation(ctx context.Context, hstReservation *rsv.HstReservation) error
	UpdateHstReservationByHstID(ctx context.Context, hstReservationID int64, hstReservation *rsv.HstReservation) error
}
