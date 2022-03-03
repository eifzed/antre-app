package order

import (
	"net/http"
	"strconv"

	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	"github.com/go-chi/chi"
)

func (h *OrderHandler) GetProductsListByShopID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	shopIDStr := chi.URLParam(r, "shopID")
	shopID, err := strconv.ParseInt(shopIDStr, 10, 64)
	if err != nil {
		err = commonerr.ErrorBadRequest("shop ID", "empty shop ID")
		commonwriter.RespondError(ctx, w, err)
		return
	}

	result, err := h.OrderUC.GetProductsListByShopID(ctx, shopID)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}
	commonwriter.RespondOKWithData(ctx, w, result)
}
