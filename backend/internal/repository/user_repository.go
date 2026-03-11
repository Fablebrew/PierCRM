package repository

import (
	"context"
	"finances/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (email, name, business_form, nalog_system, employees, business_sphere, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	return r.db.QueryRow(ctx, query, user.Email, user.Name, user.BusinessForm, user.NalogSystem, user.Employees, user.BusinessSphere, user.Password).
		Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, email, password, created_at FROM users WHERE email=$1`
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
