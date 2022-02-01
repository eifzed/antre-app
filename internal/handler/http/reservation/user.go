package reservation

import (
	"net/http"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

func (h *RsvHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.User{}
	err := bind.Bind(r, &user)
	if err != nil {
		commonwriter.RespondError(ctx, w, http.StatusBadRequest, "invalid params")
	}

}
