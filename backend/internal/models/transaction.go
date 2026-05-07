package models

import "time"

type Transaction struct {
	ID          int64     `json:"id" db:"id"`
	ProjectID   int64     `json:"project_id" db:"project_id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Type        string    `json:"type" db:"type"` // "income" или "expense"
	Account     string    `json:"account" db:"account"`
	Amount      string    `json:"amount" db:"amount"`
	IsPaid      string    `json:"is_paid" db:"is_paid"`
	Date        string    `json:"date" db:"date"`
	Responsible string    `json:"responsible" db:"responsible"`
	Comment     string    `json:"comment" db:"comment"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateTransactionRequest struct {
	ProjectID   int64  `json:"project_id" binding:"required"`
	Type        string `json:"type" binding:"required"` // "income" или "expense"
	Account     string `json:"account"`
	Amount      string `json:"amount" binding:"required"`
	IsPaid      string `json:"is_paid"`
	Date        string `json:"date"`
	Responsible string `json:"responsible"`
	Comment     string `json:"comment"`
}
