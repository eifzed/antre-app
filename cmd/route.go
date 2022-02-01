package main

import (
	"net/http"

	"github.com/eifzed/antre-app/lib/utility/urlpath"
	"github.com/go-chi/chi"
)

func getRoute(m *modules) *chi.Mux {
	router := chi.NewRouter()
	path := urlpath.New("")
	router.Route("/v1", func(v1 chi.Router) {
		v1.Group(func(auth chi.Router) {
			path.Group("/auth", func(authRoute urlpath.Routes) {
				auth.Post("/register", m.httpHandler.AuthHandler.RegisterNewAccount)
			})
		})
		v1.Group(func(antre chi.Router) {
			antre.Use(authenticate)

			path.Group("/reservations", func(reservationRoute urlpath.Routes) {
				antre.Get("/{id}", m.httpHandler.ReservationHandler.GetReservationByID)
			})
		})

	})

	// user
	return router
}
func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// TODO: get user detail from context
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
