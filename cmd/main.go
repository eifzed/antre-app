package main

import (
	"log"
	"net/http"

	rsvUC "github.com/eifzed/antre-app/internal/usecase/reservation"

	"github.com/eifzed/antre-app/internal/handler"
	rsvHandler "github.com/eifzed/antre-app/internal/handler/http/reservation"
	_ "github.com/lib/pq"
)

func main() {
	masterInfo := "postgresql://postgres:@127.0.0.1:5432/antre_app?sslmode=disable"
	slaveInfo := "postgresql://postgres:@127.0.0.1:5432/antre_app?sslmode=disable"
	conn, err := createDatabaseConnection(masterInfo, slaveInfo)
	if err != nil {
		log.Fatal(err)
	}

	reservationUC := rsvUC.NewReservationUC(&rsvUC.ReservationUC{ReservationDB: conn})
	reservationHandler := rsvHandler.NewReservationHandler(&rsvHandler.RsvHandler{ReservationUC: reservationUC})
	handler := handler.HttpHandler{ReservationHandler: reservationHandler}
	modules := newModules(modules{httpHandler: handler})
	router := getRoute(modules)
	log.Fatal(http.ListenAndServe(":9999", router))
}
