package main

import (
	"github.com/eifzed/antre-app/internal/handler"
)

type modules struct {
	httpHandler handler.HttpHandler
}

func newModules(mod modules) *modules {
	return &mod
}
