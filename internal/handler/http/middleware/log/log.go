package log

import (
	"context"
	"errors"
	"net/http"

	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/commonwriter"
	"github.com/eifzed/antre-app/lib/utility/jwt"
	"github.com/go-chi/chi"
)

type LogModule struct {
}

func NewLogModule(module *LogModule) *LogModule {
	return module
}

func (lm *LogModule) LogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rCtx := chi.RouteContext(r.Context())
		if rCtx == nil {
			logHandlerError(ctx, rw, r, errors.New("context is not Chi context"))
			return
		}

		// route := fmt.Sprintf("%s %s", rCtx.RouteMethod, rCtx.RoutePattern())

		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}

func logHandlerError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case jwt.ErrInvalid:
		err := commonerr.ErrorUnauthorized(err.Error())
		commonwriter.RespondError(ctx, w, err)
	case jwt.ErrExpired:
		err := commonerr.ErrorUnauthorized(err.Error())
		commonwriter.RespondError(ctx, w, err)
	case jwt.ErrForbidden:
		err := commonerr.ErrorForbidden(err.Error())
		commonwriter.RespondError(ctx, w, err)
	default:
		commonwriter.RespondError(ctx, w, err)
	}
}
