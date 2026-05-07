package service

import (
	"errors"
	"finances/internal/models"
	"finances/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

type RegisterInput struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	BusinessForm   string `json:"business_form"`
	NalogSystem    string `json:"nalog_system"`
	Employees      string `json:"employees"`
	BusinessSphere string `json:"business_sphere"`
	Password       string `json:"password"`
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Убрали context.Context из параметров
func (s *UserService) Register(input RegisterInput) (*models.UserResponse, error) {
	// Проверяем, не занят ли email
	existing, _ := s.repo.GetByEmail(input.Email)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("email already exists")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:          input.Email,
		Name:           input.Name,
		BusinessForm:   input.BusinessForm,
		NalogSystem:    input.NalogSystem,
		Employees:      input.Employees,
		BusinessSphere: input.BusinessSphere,
		PasswordHash:   string(hashedPassword), // Правильное поле: PasswordHash, не Password
	}

	// Убрали context.Context из вызова Create
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	// Возвращаем UserResponse
	return &models.UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		BusinessForm:   user.BusinessForm,
		BusinessSphere: user.BusinessSphere,
		Employees:      user.Employees,
		NalogSystem:    user.NalogSystem,
		CreatedAt:      user.CreatedAt,
	}, nil
}
