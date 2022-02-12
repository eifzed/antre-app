package order

import (
	"net/http"

	"github.com/eifzed/antre-app/internal/entity/order"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

func (h *OrderHandler) RegisterShop(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := order.ShopRegistration{}

	err := bind.Bind(r, &data)
	if err != nil {
		err := commonerr.ErrorBadRequest("registration params", "invalid registration params")
		commonwriter.RespondError(ctx, w, err)
	}
	err = h.OrderUC.RegisterShop(ctx, data)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}
	commonwriter.RespondOK(ctx, w)

}
