package models

type UserTicket struct {
	UserID   int `json:"user_id" db:"user_id"`
	TicketID int `json:"ticket_id" db:"ticket_id"`
}
