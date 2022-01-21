package handler

import "net/http"

type HttpHandler struct {
	ReservationHandler rsvHandler
}

type rsvHandler interface {
	GetReservationByID(w http.ResponseWriter, r *http.Request)
}
