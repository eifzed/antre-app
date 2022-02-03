package antre

import (
	"net/http"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/internal/entity/usecase/antre"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
)

type AntreHandler struct {
	AntreUC antre.AntreUCInterface
	Config  *config.Config
}

func NewAntreHandler(antreHandler *AntreHandler) *AntreHandler {
	return antreHandler
}

func (h *AntreHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.User{}
	err := bind.Bind(r, &user)
	if err != nil {
		commonwriter.RespondError(ctx, w, http.StatusBadRequest, "invalid params")
	}
}
