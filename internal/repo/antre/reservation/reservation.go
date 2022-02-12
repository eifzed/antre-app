package reservation

import (
	"context"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
)

var (
	getSession = db.GetDBSession
)

type Conn struct {
	DB *db.Connection
	// Gocrypt *gocrypt.Option
}

func NewDBConnection(conn *db.Connection) *Conn {
	return &Conn{
		DB: conn,
	}
}

func (con *Conn) GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	data := &rsv.TrxReservation{}
	session := con.DB.Slave.Context(ctx).Table("ant_trx_reservation")
	_, err := session.Where("reservation_id = ?", rsvID).Get(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (con *Conn) InsertTrxReservation(ctx context.Context, reservation *rsv.TrxReservation) error {
	session := getSession(ctx)
	if session != nil {
		session = con.DB.Slave.Context(ctx)
	}
	count, err := session.Table("ant_trx_reservation").Insert(&reservation)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}

func (con *Conn) UpdateTrxReservationByID(ctx context.Context, rsvID int64, reservation *rsv.TrxReservation) error {
	session := getSession(ctx)
	if session != nil {
		session = con.DB.Slave.Context(ctx)
	}
	count, err := session.Table("ant_trx_reservation").
		Where("reservation_id = ?", rsvID).
		Update(&reservation)
	if err != nil {
		return err
	}
	if count == 0 {
		return databaseerr.ErrorNoInsert
	}
	return nil
}
