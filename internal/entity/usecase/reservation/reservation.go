package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
)

type Reservation interface {
	reservation
}

type reservation interface {
	GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error)
	RegisterReservation(ctx context.Context, resrvation *rsv.TrxReservation) error
}
