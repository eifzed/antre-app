package main

import (
	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/handler"
)

type modules struct {
	httpHandler *handler.HttpHandler
	Config      *config.Config
}

func newModules(mod modules) *modules {
	return &mod
}
