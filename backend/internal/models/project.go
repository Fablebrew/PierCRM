package models

import "time"

type Project struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Priority    string    `json:"priority" db:"priority"`
	Stage       string    `json:"stage" db:"stage"`
	Start       string    `json:"start" db:"start_date"` // изменил db tag
	Deadline    string    `json:"deadline" db:"deadline"`
	Cost        string    `json:"cost" db:"cost"`
	Paid        string    `json:"paid" db:"paid"`
	Remainder   string    `json:"remainder" db:"remainder"` // остаток
	OtherIncome string    `json:"other_income" db:"other_income"`
	Revenue     string    `json:"revenue" db:"revenue"` // выручка
	Expense     string    `json:"expense" db:"expense"`
	Profit      string    `json:"profit" db:"profit"`
	Description string    `json:"description" db:"description"`
	Customer    string    `json:"customer" db:"customer"`
	Executor    string    `json:"executor" db:"executor"`
	TaxRate     string    `json:"tax_rate" db:"tax_rate"`
	Account     string    `json:"account" db:"account"`
	EndDate     string    `json:"end_date" db:"end_date"`
	DaysLeft    int       `json:"days_left"` // до дедлайна (вычисляемое)
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Priority    string `json:"priority"`
	Stage       string `json:"stage"`
	Start       string `json:"start"`
	Deadline    string `json:"deadline"`
	Cost        string `json:"cost"`
	Paid        string `json:"paid"`
	Remainder   string `json:"remainder"`
	OtherIncome string `json:"other_income"`
	Revenue     string `json:"revenue"`
	Expense     string `json:"expense"`
	Profit      string `json:"profit"`
	Description string `json:"description"`
	Customer    string `json:"customer"`
	Executor    string `json:"executor"`
	TaxRate     string `json:"tax_rate"`
	Account     string `json:"account"`
	EndDate     string `json:"end_date"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Priority    string `json:"priority"`
	Stage       string `json:"stage"`
	Start       string `json:"start"`
	Deadline    string `json:"deadline"`
	Cost        string `json:"cost"`
	Paid        string `json:"paid"`
	Remainder   string `json:"remainder"`
	OtherIncome string `json:"other_income"`
	Revenue     string `json:"revenue"`
	Expense     string `json:"expense"`
	Profit      string `json:"profit"`
	Description string `json:"description"`
	Customer    string `json:"customer"`
	Executor    string `json:"executor"`
	TaxRate     string `json:"tax_rate"`
	Account     string `json:"account"`
	EndDate     string `json:"end_date"`
}

type ProjectFilter struct {
	SortBy    string `json:"sort_by"`    // priority, stage, created_at, deadline
	Order     string `json:"order"`      // asc, desc
	Stage     string `json:"stage"`      // фильтр по этапу
	Priority  string `json:"priority"`   // фильтр по приоритету
	Search    string `json:"search"`     // поиск по названию
	StartDate string `json:"start_date"` // период с
	EndDate   string `json:"end_date"`   // период по
}
