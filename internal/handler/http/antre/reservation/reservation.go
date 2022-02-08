package reservation

import (
	"net/http"
	"strconv"

	"github.com/eifzed/antre-app/internal/entity/reservation"
	rsvUC "github.com/eifzed/antre-app/internal/entity/usecase/antre/reservation"
	"github.com/go-chi/chi"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

type RsvHandler struct {
	ReservationUC rsvUC.ReservationUCInterface
	Config        *config.Config
}

func NewReservationHandler(rsvHandler *RsvHandler) *RsvHandler {
	return rsvHandler
}

func (h *RsvHandler) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsvIDParams := chi.URLParam(r, "id")
	rsvID, err := strconv.ParseInt(rsvIDParams, 10, 64)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}

	result, err := h.ReservationUC.GetReservationByID(ctx, rsvID)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	commonwriter.RespondOKWithData(ctx, w, result)

}

func (h *RsvHandler) RegisterReservation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &reservation.TrxReservation{}

	err := bind.Bind(r, &data)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
	}
	err = h.ReservationUC.RegisterReservation(ctx, data)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	commonwriter.RespondOKWithData(ctx, w, data)

}
