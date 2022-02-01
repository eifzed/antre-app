package main

import (
	"log"
	"net/http"

	"github.com/eifzed/antre-app/internal/config"
	rsvUC "github.com/eifzed/antre-app/internal/usecase/reservation"

	"github.com/eifzed/antre-app/internal/handler"
	rsvHandler "github.com/eifzed/antre-app/internal/handler/http/reservation"
	_ "github.com/lib/pq"
)

func main() {
	secrete := config.GetSecretes()
	if secrete == nil {
		log.Fatal("failed to get secretes")
		return
	}
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg.Secretes = secrete
	conn, err := createDatabaseConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err)
	}
	reservationUC := rsvUC.NewReservationUC(&rsvUC.ReservationUC{
		ReservationDB: conn,
		Config:        cfg,
	})
	reservationHandler := rsvHandler.NewReservationHandler(&rsvHandler.RsvHandler{
		ReservationUC: reservationUC,
		Config:        cfg,
	})
	handler := handler.HttpHandler{ReservationHandler: reservationHandler}
	modules := newModules(modules{
		httpHandler: &handler,
		Config:      cfg,
	})
	router := getRoute(modules)
	log.Fatal(http.ListenAndServe(cfg.Server.HTTP.Address, router))
}
