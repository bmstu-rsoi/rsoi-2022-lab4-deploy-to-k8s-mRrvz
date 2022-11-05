package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Flight struct {
	ID            int    `json:"id"`
	FlightNumber  string `json:"flightNumber"`
	Date          string `json:"date"`
	FromAirportId int    `json:"fromAirportId"`
	ToAirportId   int    `json:"toAirportId"`
	Price         int    `json:"price"`
}

type Airport struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func (p *Flight) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.FlightNumber, validation.Required),
		validation.Field(&p.Date, validation.Required, validation.Date("2006-01-01")),
		validation.Field(&p.FromAirportId, validation.Required, validation.Min(0)),
		validation.Field(&p.ToAirportId, validation.Required, validation.Min(0)),
		validation.Field(&p.Price, validation.Required, validation.Min(0)),
	)
}

func (p *Airport) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.City, validation.Required),
		validation.Field(&p.Country, validation.Required),
	)
}
