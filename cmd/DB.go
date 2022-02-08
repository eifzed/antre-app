package main

import (
	"log"

	db "github.com/eifzed/antre-app/lib/database/xorm"
)

func createDBConnection(masterInfo string, slaveInfo string) (*db.Connection, error) {
	masterConn, err := db.ConnetDB(masterInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	slaveConn, err := db.ConnetDB(slaveInfo)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &db.Connection{Master: masterConn, Slave: slaveConn}, nil
}
