package main

import (
	"github.com/eifzed/antre-app/internal/handler"
)

type modules struct {
	handler handler.Handler
}

func newModules(mod modules) *modules {
	return &mod
}
