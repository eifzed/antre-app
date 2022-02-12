package handler

import "net/http"

type rsvHandlerInterface interface {
	// reservation handler
	GetReservationByID(w http.ResponseWriter, r *http.Request)
	RegisterReservation(w http.ResponseWriter, r *http.Request)

	// shop handler
	RegisterShop(w http.ResponseWriter, r *http.Request)
}
