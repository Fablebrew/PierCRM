package repository

import (
	"context"
	"finances/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
        INSERT INTO users (name, email, password_hash, business_form, business_sphere, employees, nalog_system, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at
    `

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.BusinessForm,
		user.BusinessSphere,
		user.Employees,
		user.NalogSystem,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

// GetByEmail без context.Context параметра
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, email, password_hash, business_form, business_sphere, employees, nalog_system, created_at, updated_at
        FROM users
        WHERE email = $1
    `

	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.BusinessForm,
		&user.BusinessSphere,
		&user.Employees,
		&user.NalogSystem,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	user := &models.User{}
	query := `
        SELECT id, name, email, password_hash, business_form, business_sphere, employees, nalog_system, created_at, updated_at
        FROM users
        WHERE id = $1
    `

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.BusinessForm,
		&user.BusinessSphere,
		&user.Employees,
		&user.NalogSystem,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
