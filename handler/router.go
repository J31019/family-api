package handler

import (
	"net/http"

	"github.com/MonkeyBuisness/family-api/controller"
	"github.com/gorilla/mux"
)

// API represents gorilla mux driver and invite with controller functions.
type API struct {
	*mux.Router
	service controller.Service
}

// New creates a router for this microservice.
func New(service controller.Service) API {
	api := API{
		Router:  mux.NewRouter(),
		service: service,
	}
	api.Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(api.login)
	return api
}
