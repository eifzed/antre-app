package reservation

import (
	"context"

	rsvRepo "github.com/eifzed/antre-app/internal/entity/repo/reservation"
	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
)

type ReservationUC struct {
	ReservationDB rsvRepo.Reservation
}

func NewReservationUC(resrvation *ReservationUC) *ReservationUC {
	return resrvation
}

func (uc *ReservationUC) GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	return uc.ReservationDB.GetReservationByID(ctx, rsvID)
}
