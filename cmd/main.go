package main

import (
	"log"

	rsvRepo "github.com/eifzed/antre-app/internal/entity/repo/reservation"

	"github.com/eifzed/antre-app/internal/handler"
	rsvHandler "github.com/eifzed/antre-app/internal/handler/http/reservation"
	_ "github.com/lib/pq"
)

var conn *rsvRepo.Conn

func main() {
	masterConn, err := rsvRepo.ConnetDB("postgresql://postgres:@127.0.0.1:5432/antre-app?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	slaveConn, err := rsvRepo.ConnetDB("postgresql://postgres:@127.0.0.1:5432/antre-app?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	reservationHandler := rsvHandler.NewReservationHandler()
	handler := handler.Handler{ReservationHandler: reservationHandler}
	modules := newModules(modules{handler: handler})
	conn = rsvRepo.NewDBConnection(&rsvRepo.Connection{Master: masterConn, Slave: slaveConn})
	getRoute(modules)
}
