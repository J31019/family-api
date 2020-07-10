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
	api.Path("/register").
		Methods(http.MethodPost).
		HandlerFunc(api.register)
	api.Path("/submit").
		Methods(http.MethodPost).
		HandlerFunc(api.submit)
	api.Path("/video").
		Methods(http.MethodGet).
		HandlerFunc(api.videoList)
	api.Path("/video/{id}").
		Methods(http.MethodGet).
		HandlerFunc(api.videoByID)
	return api
}
