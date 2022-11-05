package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/flights", GetAllFlightHandler).Methods("GET")
	router.HandleFunc("/api/v1/flight/{flightNumber}", GetFlightHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/flight/airport/{airportId}", GetAirportHandler).Methods("GET", "OPTIONS")

	return router
}
