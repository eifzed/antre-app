package antre

import (
	"net/http"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/internal/entity/usecase/antre"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	bind "github.com/eifzed/antre-app/lib/common/handler"
	"github.com/go-chi/chi"
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

	user := auth.UserDetail{}
	err := bind.Bind(r, &user)
	if err != nil {
		newErr := commonerr.ErrorBadRequest("Params", "Invalid Params")
		commonwriter.RespondError(ctx, w, newErr)
		return
	}
	resp, err := h.AntreUC.RegisterNewAccount(ctx, user)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}

	commonwriter.RespondOKWithData(ctx, w, resp)
}

func (h *AntreHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := auth.LoginParams{}
	err := bind.Bind(r, &user)
	if err != nil {
		newErr := commonerr.ErrorBadRequest("Params", "Invalid Params")
		commonwriter.RespondError(ctx, w, newErr)
		return
	}
	resp, err := h.AntreUC.Login(ctx, user)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}

	commonwriter.RespondOKWithData(ctx, w, resp)
}

func (h *AntreHandler) AssignNewRoleToUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	role := chi.URLParam(r, "role")
	if role == "" {
		newErr := commonerr.ErrorBadRequest("Role", "Empty Role")
		commonwriter.RespondError(ctx, w, newErr)
		return
	}
	err := h.AntreUC.AssignNewRoleToUser(ctx, role)
	if err != nil {
		commonwriter.RespondError(ctx, w, err)
		return
	}

	commonwriter.RespondOK(ctx, w)
}
