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
		v1.Group(func(user chi.Router) {
			path.Group("/user", func(userRoute urlpath.Routes) {
				user.Post("/register", m.httpHandler.AntreHandler.RegisterNewAccount)
			})
		})
		v1.Group(func(antre chi.Router) {
			antre.Use(m.AuthModule.AuthHandler)
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
