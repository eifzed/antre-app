package reservation

import (
	"context"

	db "github.com/eifzed/antre-app/internal/entity/database"
	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
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
