package reservation

import (
	"context"
	"encoding/json"
	"net/http"

	rsvUC "github.com/eifzed/antre-app/internal/entity/usecase/reservation"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
)

type RsvHandler struct {
	ReservationUC rsvUC.Reservation
}

func NewReservationHandler(rsvHandler *RsvHandler) *RsvHandler {
	return rsvHandler
}

func (h *RsvHandler) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.ReservationUC.GetReservationByID(ctx, 123)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)

}

func GetReservationByIDUC(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	result, err := GetReservationByIDRepo(ctx, rsvID)
	if err != nil {
		return nil, err
	}
	return result, err

}

func GetReservationByIDRepo(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	result := &rsv.TrxReservation{}

	// has, err := conn.DB.Slave.Context(ctx).
	// 	Table("ant_trx_reservation").
	// 	Where("reservation_id = ?", rsvID).
	// 	Get(result)

	// if err != nil {
	// 	return nil, err
	// }

	// if !has {
	// 	return nil, errors.New("not found")
	// }
	return result, nil
}
