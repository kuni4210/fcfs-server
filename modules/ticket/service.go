package ticket

import (
	"database/sql"
)

type TicketService struct {
	postgres *sql.DB
}

func NewTicketService(postgres *sql.DB) *TicketService {
	return &TicketService{postgres: postgres}
}
