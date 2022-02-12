package handler

import "net/http"

type HttpHandler struct {
	OrderHandler rsvHandlerInterface
	AntreHandler antreHandlerInterface
}

type antreHandlerInterface interface {
	RegisterNewAccount(w http.ResponseWriter, r *http.Request)
	AssignNewRoleToUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type AuthModuleInterface interface {
	AuthHandler(next http.Handler) http.Handler
}
