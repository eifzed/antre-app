package main

import (
	"log"

	rsvDB "github.com/eifzed/antre-app/internal/entity/database"
	rsvRepo "github.com/eifzed/antre-app/internal/repo/reservation"
)

func createDatabaseConnection(masterInfo string, slaveInfo string) (*rsvRepo.Conn, error) {
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
