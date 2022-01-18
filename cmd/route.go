package main

import "github.com/gorilla/mux"

func getRoute(m *modules) *mux.Router {
	router := newRouter().PathPrefix("/v1")

	antreRouter := router.PathPrefix("/antre").Subrouter()
	antreRouter.HandleFunc("/reservations", m.handler.ReservationHandler.GetReservationByID).Methods("GET")
	return antreRouter
}
func newRouter() *mux.Router {
	r := mux.NewRouter()
	//TODO:Add Middleware
	return r
}
