package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (conn *Conn) GetHstReservationByHstID(ctx context.Context, hstReservationID int64) (*rsv.HstReservation, error) {
	hstReservation := &rsv.HstReservation{}
	session := conn.DB.Master.Context(ctx).Table("ant_hst_reservation")
	has, err := session.Where("hst_reservation_id = ?", hstReservationID).Get(hstReservation)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return hstReservation, nil
}

func (conn *Conn) GetHstReservationByRsvID(ctx context.Context, rsvID int64) (*rsv.HstReservation, error) {
	hstReservation := &rsv.HstReservation{}
	session := conn.DB.Master.Context(ctx).Table("ant_hst_reservation")
	has, err := session.Where("reservation_id = ?", rsvID).Get(hstReservation)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, databaseerr.ErrorDataNotFound
	}
	return hstReservation, nil
}

func (conn *Conn) InsertHstReservation(ctx context.Context, hstReservation *rsv.HstReservation) error {
	session := conn.DB.Master.Context(ctx).Table("ant_hst_reservation")
	count, err := session.Insert(hstReservation)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (conn *Conn) UpdateHstReservationByHstID(ctx context.Context, hstReservationID int64, hstReservation *rsv.HstReservation) error {
	session := conn.DB.Master.Context(ctx).Table("ant_hst_reservation")
	count, err := session.Where("hst_reservation_id = ?", hstReservationID).Update(hstReservation)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoUpdate
	}
	return nil
}
