package repository

import (
	"context"
	"finances/internal/models"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(userID int64, req *models.CreateTransactionRequest) (*models.Transaction, error) {
	query := `
		INSERT INTO transactions (project_id, user_id, type, account, amount, is_paid, date, responsible, comment, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING *
	`

	now := time.Now()
	t := &models.Transaction{}

	err := r.db.QueryRow(context.Background(), query,
		req.ProjectID, userID, req.Type, req.Account, req.Amount, req.IsPaid, req.Date, req.Responsible, req.Comment,
		now, now,
	).Scan(
		&t.ID, &t.ProjectID, &t.UserID, &t.Type, &t.Account, &t.Amount, &t.IsPaid, &t.Date, &t.Responsible, &t.Comment,
		&t.CreatedAt, &t.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return t, nil
}

func (r *TransactionRepository) GetByProjectID(userID, projectID int64) ([]models.Transaction, error) {
	query := `SELECT * FROM transactions WHERE user_id = $1 AND project_id = $2 ORDER BY created_at DESC`

	rows, err := r.db.Query(context.Background(), query, userID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(
			&t.ID, &t.ProjectID, &t.UserID, &t.Type, &t.Account, &t.Amount, &t.IsPaid, &t.Date, &t.Responsible, &t.Comment,
			&t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	if transactions == nil {
		transactions = []models.Transaction{}
	}
	return transactions, nil
}

// GetTotals возвращает суммы доходов и расходов по проекту
func (r *TransactionRepository) GetTotals(userID, projectID int64) (totalIncome, totalExpense float64, err error) {
	query := `
		SELECT 
			COALESCE(SUM(CASE WHEN type = 'income' THEN CAST(amount AS DECIMAL) ELSE 0 END), 0),
			COALESCE(SUM(CASE WHEN type = 'expense' THEN CAST(amount AS DECIMAL) ELSE 0 END), 0)
		FROM transactions 
		WHERE user_id = $1 AND project_id = $2
	`

	err = r.db.QueryRow(context.Background(), query, userID, projectID).Scan(&totalIncome, &totalExpense)
	return
}
