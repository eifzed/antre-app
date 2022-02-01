package handler

import "net/http"

type HttpHandler struct {
	ReservationHandler rsvHandler
	AuthHandler        authHandler
}

type rsvHandler interface {
	GetReservationByID(w http.ResponseWriter, r *http.Request)
}

type authHandler interface {
	RegisterNewAccount(w http.ResponseWriter, r *http.Request)
}
