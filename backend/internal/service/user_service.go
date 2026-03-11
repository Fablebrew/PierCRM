package service

import (
	"context"
	"finances/internal/models"
	"finances/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}
type RegisterInput struct {
	Email          string
	Name           string
	BusinessForm   string
	NalogSystem    string
	Employees      string
	BusinessSphere string
	Password       string
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, input RegisterInput) (*models.User, error) {
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
		Password:       string(hashedPassword),
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
