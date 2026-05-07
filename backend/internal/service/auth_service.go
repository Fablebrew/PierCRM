package service

import (
	"errors"
	"finances/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	secret   string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{
		userRepo: repo,
		secret:   secret,
	}
}

// Login проверяет пользователя и возвращает JWT
func (s *AuthService) Login(email, password string) (string, error) {
	// Убрали context.Context из вызова GetByEmail
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Используем PasswordHash вместо Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// Создаём JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.secret))
}
