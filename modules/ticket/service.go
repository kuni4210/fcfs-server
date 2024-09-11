package ticket

import (
	"database/sql"
	"fcfs-server/config"
)

type TicketService struct {
	postgres *sql.DB
	cfg      *config.Config
}

func NewTicketService(postgres *sql.DB, cfg *config.Config) *TicketService {
	return &TicketService{postgres: postgres, cfg: cfg}
}
