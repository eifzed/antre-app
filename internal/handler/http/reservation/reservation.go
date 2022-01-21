package reservation

import (
	"net/http"
	"strconv"

	rsvUC "github.com/eifzed/antre-app/internal/entity/usecase/reservation"

	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

type RsvHandler struct {
	ReservationUC rsvUC.Reservation
}

func NewReservationHandler(rsvHandler *RsvHandler) *RsvHandler {
	return rsvHandler
}

func (h *RsvHandler) GetReservationByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsvIDParams := bind.GetURLParam(r, "id")
	rsvID, err := strconv.ParseInt(rsvIDParams, 10, 64)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}

	result, err := h.ReservationUC.GetReservationByID(ctx, rsvID)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}
	commonwriter.RespondOKWithData(ctx, w, result)

}
