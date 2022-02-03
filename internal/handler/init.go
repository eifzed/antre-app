package handler

import "net/http"

type HttpHandler struct {
	ReservationHandler rsvHandlerInterface
	AntreHandler       antreHandlerInterface
}

type rsvHandlerInterface interface {
	GetReservationByID(w http.ResponseWriter, r *http.Request)
}

type antreHandlerInterface interface {
	RegisterNewAccount(w http.ResponseWriter, r *http.Request)
}

type AuthModuleInterface interface {
	AuthHandler(next http.Handler) http.Handler
}
