package main

import (
	"log"

	"github.com/eifzed/antre-app/internal/config"
	orderUC "github.com/eifzed/antre-app/internal/usecase/antre/order"

	"github.com/eifzed/antre-app/internal/handler"
	antreHandler "github.com/eifzed/antre-app/internal/handler/http/antre"
	orderHandler "github.com/eifzed/antre-app/internal/handler/http/antre/order"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	antreRepo "github.com/eifzed/antre-app/internal/repo/antre"
	orderRepo "github.com/eifzed/antre-app/internal/repo/antre/order"
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
	orderConn, err := createDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	antreConn, err := createDBConnection(cfg.Secretes.Data.DBMaster.DSN, cfg.Secretes.Data.DBSlave.DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	antreDB := antreRepo.NewDBConnection(antreConn)
	orderDB := orderRepo.NewDBConnection(orderConn)
	orderTransaction := db.GetDBTransaction(orderDB.DB)
	antreTransaction := db.GetDBTransaction(antreDB.DB)
	antreUC := antreUC.NewAntreUC(&antreUC.AntreUC{
		AntreDB:     antreDB,
		Config:      cfg,
		Transaction: antreTransaction,
	})

	orderUC := orderUC.NewOrderUC(&orderUC.OrderUC{
		OrderDB:     orderDB,
		Config:      cfg,
		Transaction: orderTransaction,
	})

	antreHandler := antreHandler.NewAntreHandler(&antreHandler.AntreHandler{
		AntreUC: antreUC,
		Config:  cfg,
	})

	orderHandler := orderHandler.NewOrderHandler(&orderHandler.OrderHandler{
		OrderUC: orderUC,
		Config:  cfg,
	})
	handler := handler.HttpHandler{
		OrderHandler: orderHandler,
		AntreHandler: antreHandler,
	}
	authHandler := auth.NewAuthModule(&auth.AuthModule{
		JWTCertificate: cfg.Secretes.Data.JWTCertificate,
		RouteRoles:     cfg.RouteRoles,
	})
	modules := newModules(modules{
		httpHandler: &handler,
		Config:      cfg,
		AuthModule:  authHandler,
	})
	router := getRoute(modules)
	ListenAndServe(cfg.Server.HTTP.Address, router)
}
