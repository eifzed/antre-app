package main

import (
	"log"
	"net/http"

	"github.com/eifzed/antre-app/internal/config"
	rsvUC "github.com/eifzed/antre-app/internal/usecase/antre/reservation"

	"github.com/eifzed/antre-app/internal/handler"
	antreHandler "github.com/eifzed/antre-app/internal/handler/http/antre"
	rsvHandler "github.com/eifzed/antre-app/internal/handler/http/antre/reservation"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	antreRepo "github.com/eifzed/antre-app/internal/repo/antre"
	rsvRepo "github.com/eifzed/antre-app/internal/repo/antre/reservation"
	antreUC "github.com/eifzed/antre-app/internal/usecase/antre"
	db "github.com/eifzed/antre-app/lib/database/xorm"
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
	rsvConn, err := createDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err)
	}
	antreConn, err := createDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err)
	}
	antreDB := antreRepo.NewDBConnection(antreConn)
	rsvDB := rsvRepo.NewDBConnection(rsvConn)
	rsvTransaction := db.GetDBTransaction(rsvDB.DB)
	antreTransaction := db.GetDBTransaction(antreDB.DB)
	antreUC := antreUC.NewAntreUC(&antreUC.AntreUC{
		AntreDB:     antreDB,
		Config:      cfg,
		Transaction: antreTransaction,
	})

	reservationUC := rsvUC.NewReservationUC(&rsvUC.ReservationUC{
		ReservationDB: rsvDB,
		Config:        cfg,
		Transaction:   rsvTransaction,
	})

	antreHandler := antreHandler.NewAntreHandler(&antreHandler.AntreHandler{
		AntreUC: antreUC,
		Config:  cfg,
	})

	reservationHandler := rsvHandler.NewReservationHandler(&rsvHandler.RsvHandler{
		ReservationUC: reservationUC,
		Config:        cfg,
	})
	handler := handler.HttpHandler{
		ReservationHandler: reservationHandler,
		AntreHandler:       antreHandler,
	}
	authHandler := auth.NewAuthModule(&auth.AuthModule{
		JWTCertificate: cfg.Secretes.Data.JWTCertificate,
		RouteRoles:     cfg.RouteRoles, // TODO: updte route roles
	})
	modules := newModules(modules{
		httpHandler: &handler,
		Config:      cfg,
		AuthModule:  authHandler,
	})
	router := getRoute(modules)
	log.Fatal(http.ListenAndServe(cfg.Server.HTTP.Address, router))
}
