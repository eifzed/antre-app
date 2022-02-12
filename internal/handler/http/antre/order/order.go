package order

import (
	"net/http"
	"strconv"

	"github.com/eifzed/antre-app/internal/entity/order"
	rsvUC "github.com/eifzed/antre-app/internal/entity/usecase/antre/order"
	"github.com/go-chi/chi"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

type RsvHandler struct {
	OrderUC rsvUC.OrderUCInterface
	Config  *config.Config
}

func NewOrderHandler(rsvHandler *RsvHandler) *RsvHandler {
	return rsvHandler
}

func (h *RsvHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	rsvIDParams := chi.URLParam(r, "id")
	rsvID, err := strconv.ParseInt(rsvIDParams, 10, 64)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}

	result, err := h.OrderUC.GetOrderByID(ctx, rsvID)
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

func (h *RsvHandler) RegisterOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := order.OrderRegistration{}

	err := bind.Bind(r, &data)
	if err != nil {
		err = commonerr.ErrorBadRequest("invalid params", "invalid order params")
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	err = h.OrderUC.RegisterOrder(ctx, data)
	if err != nil {
		commonwriter.RespondDefaultError(ctx, w, err)
		return
	}
	commonwriter.RespondOK(ctx, w)
}
