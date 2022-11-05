package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/flight-service/internal/repository"
	"github.com/gorilla/mux"
)

func GetAllFlightHandler(w http.ResponseWriter, r *http.Request) {
	flightRepo := repository.FlightRepository{}
	flight, err := flightRepo.GetAllFlights()
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(flight); err != nil {
		log.Printf("Failed to encode flight: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetFlightHandler(w http.ResponseWriter, r *http.Request) {
	flightRepo := repository.FlightRepository{}
	params := mux.Vars(r)
	flight, err := flightRepo.GetFlightByNumber(params["flightNumber"])
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(flight); err != nil {
		log.Printf("Failed to encode flight: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetAirportHandler(w http.ResponseWriter, r *http.Request) {
	flightRepo := repository.FlightRepository{}
	params := mux.Vars(r)
	flight, err := flightRepo.GetAirportByID(params["airportId"])
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(flight); err != nil {
		log.Printf("Failed to encode flight: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
