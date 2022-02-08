package reservation

import (
	"context"

	"github.com/eifzed/antre-app/internal/config"
	rsvRepo "github.com/eifzed/antre-app/internal/entity/repo/antre/reservation"
	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	db "github.com/eifzed/antre-app/lib/database/xorm"
)

type ReservationUC struct {
	ReservationDB rsvRepo.Reservation
	Config        *config.Config
	Transaction   *db.DBTransaction
}

func NewReservationUC(resrvation *ReservationUC) *ReservationUC {
	return resrvation
}

func (uc *ReservationUC) GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	// TODO: get user detail from context
	return uc.ReservationDB.GetReservationByID(ctx, rsvID)
}

func (uc *ReservationUC) RegisterReservation(ctx context.Context, reservation *rsv.TrxReservation) error {
	// TODO: get user detail from context
	err := uc.ReservationDB.InsertTrxReservation(ctx, reservation)
	if err != nil {
		return err
	}

	hstReservation := &rsv.HstReservation{
		ReservationID:   reservation.ReservationID,
		CustomerID:      reservation.CustomerID,
		ShopID:          reservation.ShopID,
		ReservationTime: reservation.ReservationTime,
		ReservationType: reservation.ReservationType,
		StatusID:        reservation.StatusID,
		CustomerNote:    reservation.CustomerNote,
		ShopNote:        reservation.ShopNote,
		UpdaterID:       reservation.CustomerID, // TODO: get from context
		Reason:          "New Reservation",
	}
	uc.ReservationDB.InsertHstReservation(ctx, hstReservation)
	return nil
}
