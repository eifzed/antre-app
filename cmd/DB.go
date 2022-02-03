package main

import (
	"log"

	rsvDB "github.com/eifzed/antre-app/internal/entity/database"
	antreRepo "github.com/eifzed/antre-app/internal/repo/antre"
	rsvRepo "github.com/eifzed/antre-app/internal/repo/antre/reservation"
)

func createRsvDBConnection(masterInfo string, slaveInfo string) (*rsvRepo.Conn, error) {
	masterConn, err := rsvDB.ConnetDB(masterInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	slaveConn, err := rsvDB.ConnetDB(slaveInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rsvRepo.NewDBConnection(&rsvDB.Connection{Master: masterConn, Slave: slaveConn}), nil
}

func createAntreDBConnection(masterInfo string, slaveInfo string) (*antreRepo.Conn, error) {
	masterConn, err := rsvDB.ConnetDB(masterInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	slaveConn, err := rsvDB.ConnetDB(slaveInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return antreRepo.NewDBConnection(&rsvDB.Connection{Master: masterConn, Slave: slaveConn}), nil
}
