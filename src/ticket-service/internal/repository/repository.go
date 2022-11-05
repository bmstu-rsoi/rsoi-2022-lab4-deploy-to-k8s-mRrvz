package repository

import (
	"database/sql"
	"fmt"

	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/ticket-service/internal/db"
	"github.com/bmstu-rsoi/rsoi-2022-lab2-microservices-mRrvz/src/ticket-service/internal/models"
)

type Repository interface {
	GetTicketsByUsername(flightNumber string) ([]*models.Ticket, error)
	CreateTicket(*models.Ticket) error
}

type TicketRepository struct {
	db *sql.DB
}

const selectTicketsByUsername = `SELECT * FROM ticket WHERE username = $1;`

func (r *TicketRepository) GetTicketsByUsername(username string) ([]*models.Ticket, error) {
	r.db = db.CreateConnection()
	defer r.db.Close()

	var tickets []*models.Ticket
	rows, err := r.db.Query(selectTicketsByUsername, username)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %s", err)
	}

	for rows.Next() {
		ticket := new(models.Ticket)
		rows.Scan(
			&ticket.ID,
			&ticket.TicketUID,
			&ticket.Username,
			&ticket.FlightNumber,
			&ticket.Price,
			&ticket.Status)

		if err != nil {
			return nil, fmt.Errorf("failed to execute the query: %s", err)
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

const createTicket = `INSERT INTO ticket (ticket_uid, username, flight_number, price, status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

func (r *TicketRepository) CreateTicket(ticket *models.Ticket) error {
	r.db = db.CreateConnection()
	defer r.db.Close()

	_, err := r.db.Query(
		createTicket,
		ticket.TicketUID,
		ticket.Username,
		ticket.FlightNumber,
		ticket.Price,
		ticket.Status,
	)

	return err
}
