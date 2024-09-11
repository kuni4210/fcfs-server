package auth

import (
	"database/sql"
	"fcfs-server/config"
	"fcfs-server/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	postgres *sql.DB
	cfg      *config.Config
}

func NewAuthService(postgres *sql.DB, cfg *config.Config) *AuthService {
	return &AuthService{postgres: postgres, cfg: cfg}
}

func (as *AuthService) Login(username, password string) (string, error) {
	var hashedPassword string
	var userID string
	err := as.postgres.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&userID, &hashedPassword)
	if err != nil {
		return "", fmt.Errorf("invalid credentials: %w", err)
	}

	if !as.checkPassword(password, hashedPassword) {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := as.generateToken(userID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (as *AuthService) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(as.cfg.Jwt.SecretKey))
}

func (as *AuthService) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(as.cfg.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func (as *AuthService) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	err := as.postgres.QueryRow("SELECT id, username FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (as *AuthService) checkPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
