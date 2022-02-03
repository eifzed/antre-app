package main

import (
	"log"
	"net/http"

	"github.com/eifzed/antre-app/internal/config"
	rsvUC "github.com/eifzed/antre-app/internal/usecase/antre/reservation"
	"github.com/eifzed/antre-app/lib/utility/jwt"

	"github.com/eifzed/antre-app/internal/handler"
	antreHandler "github.com/eifzed/antre-app/internal/handler/http/antre"
	rsvHandler "github.com/eifzed/antre-app/internal/handler/http/antre/reservation"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	antreUC "github.com/eifzed/antre-app/internal/usecase/antre"
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
	rsvConn, err := createRsvDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err)
	}
	antreConn, err := createAntreDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err)
	}
	antreUC := antreUC.NewAntreUC(&antreUC.AntreUC{
		AntreDB: antreConn,
		Config:  cfg,
	})

	reservationUC := rsvUC.NewReservationUC(&rsvUC.ReservationUC{
		ReservationDB: rsvConn,
		Config:        cfg,
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
		RouteRoles:     map[string]*jwt.RouteRoles{}, // TODO: updte route roles
	})
	modules := newModules(modules{
		httpHandler: &handler,
		Config:      cfg,
		AuthHandler: authHandler,
	})
	router := getRoute(modules)
	log.Fatal(http.ListenAndServe(cfg.Server.HTTP.Address, router))
}
