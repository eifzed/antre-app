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
	return nil, nil
}
