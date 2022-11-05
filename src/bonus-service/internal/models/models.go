package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Privilege struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
	Balance  int    `json:"balance"`
}

type PrivilegeHistory struct {
	ID            int    `json:"id"`
	PrivilegeID   int    `json:"privilegeId"`
	TicketUID     string `json:"ticketUid"`
	Date          string `json:"date"`
	BalanceDiff   int    `json:"balanceDiff"`
	OperationType string `json:"operationType"`
}

func (p *Privilege) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Username, validation.Required),
		validation.Field(&p.Status, validation.Required, validation.In("BRONZE", "SILVER", "GOLD")),
		validation.Field(&p.Balance, validation.Required, validation.Min(0)),
	)
}
