package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Ticket struct {
	ID           int    `json:"id"`
	TicketUID    string `json:"ticketUid"`
	Username     string `json:"username"`
	FlightNumber string `json:"flightNumber"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
}

func (p *Ticket) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.TicketUID, is.UUID),
		validation.Field(&p.Username, validation.Required),
		validation.Field(&p.FlightNumber, validation.Required),
		validation.Field(&p.Price, validation.Required, validation.Min(0)),
		validation.Field(&p.Status, validation.Required),
	)
}
