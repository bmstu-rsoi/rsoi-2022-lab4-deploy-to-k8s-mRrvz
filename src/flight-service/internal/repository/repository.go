package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/flight-service/internal/db"
	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/flight-service/internal/models"
)

type Repository interface {
	GetAllFlights() (*models.Flight, error)
	GetFlightByNumber(flightNumber string) (*models.Flight, error)
	GetAirportByID(airportID string) (*models.Airport, error)
}

type FlightRepository struct {
	db *sql.DB
}

const selectAllFlights = `SELECT * FROM flight;`

func (r *FlightRepository) GetAllFlights() ([]*models.Flight, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var flights []*models.Flight
	rows, err := r.db.Query(selectAllFlights)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		f := new(models.Flight)
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.Date, &f.FromAirportId, &f.ToAirportId, &f.Price); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		flights = append(flights, f)
	}

	defer rows.Close()

	return flights, nil
}

const selectFlightByNumber = `SELECT * FROM flight WHERE flight_number = $1;`

func (r *FlightRepository) GetFlightByNumber(flightNumber string) (*models.Flight, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var flight models.Flight
	row := r.db.QueryRow(selectFlightByNumber, flightNumber)
	err := row.Scan(
		&flight.ID,
		&flight.FlightNumber,
		&flight.Date,
		&flight.FromAirportId,
		&flight.ToAirportId,
		&flight.Price,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &flight, err
		}
	}

	return &flight, nil
}

const selectAirportByID = `SELECT * FROM airport WHERE id = $1;`

func (r *FlightRepository) GetAirportByID(airportID string) (*models.Airport, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var airport models.Airport
	row := r.db.QueryRow(selectAirportByID, airportID)
	err := row.Scan(
		&airport.ID,
		&airport.Name,
		&airport.City,
		&airport.Country,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &airport, err
		}
	}

	return &airport, nil
}
