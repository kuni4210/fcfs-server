package auth

import (
	"database/sql"
)

type AuthService struct {
	postgres  *sql.DB
	jwtSecret string
}

func NewAuthService(postgres *sql.DB, jwtSecret string) *AuthService {
	return &AuthService{postgres: postgres, jwtSecret: jwtSecret}
}
