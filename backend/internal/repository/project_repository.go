package repository

import (
	"context"
	"finances/internal/models"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(userID int64, req *models.CreateProjectRequest) (*models.Project, error) {
	query := `
		INSERT INTO projects (user_id, name, priority, stage, start_date, deadline, cost, paid, remainder, other_income, revenue, expense, profit, description, customer, executor, tax_rate, account, end_date, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21)
		RETURNING *
	`

	now := time.Now()
	p := &models.Project{}

	err := r.db.QueryRow(context.Background(), query,
		userID, req.Name, req.Priority, req.Stage, req.Start, req.Deadline,
		req.Cost, req.Paid, req.Remainder, req.OtherIncome, req.Revenue, req.Expense, req.Profit,
		req.Description, req.Customer, req.Executor, req.TaxRate, req.Account, req.EndDate,
		now, now,
	).Scan(
		&p.ID, &p.UserID, &p.Name, &p.Priority, &p.Stage, &p.Start, &p.Deadline,
		&p.Cost, &p.Paid, &p.Remainder, &p.OtherIncome, &p.Revenue, &p.Expense, &p.Profit,
		&p.Description, &p.Customer, &p.Executor, &p.TaxRate, &p.Account, &p.EndDate,
		&p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	calculateProject(p) // ← Вычисляем производные поля
	return p, nil
}

func (r *ProjectRepository) GetAll(userID int64) ([]models.Project, error) {
	query := `SELECT * FROM projects WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.Name, &p.Priority, &p.Stage, &p.Start, &p.Deadline,
			&p.Cost, &p.Paid, &p.Remainder, &p.OtherIncome, &p.Revenue, &p.Expense, &p.Profit,
			&p.Description, &p.Customer, &p.Executor, &p.TaxRate, &p.Account, &p.EndDate,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		calculateProject(&p) // ← Вычисляем для каждого проекта
		projects = append(projects, p)
	}

	if projects == nil {
		projects = []models.Project{}
	}
	return projects, nil
}

func (r *ProjectRepository) GetByID(userID, id int64) (*models.Project, error) {
	query := `SELECT * FROM projects WHERE id = $1 AND user_id = $2`
	p := &models.Project{}

	err := r.db.QueryRow(context.Background(), query, id, userID).Scan(
		&p.ID, &p.UserID, &p.Name, &p.Priority, &p.Stage, &p.Start, &p.Deadline,
		&p.Cost, &p.Paid, &p.Remainder, &p.OtherIncome, &p.Revenue, &p.Expense, &p.Profit,
		&p.Description, &p.Customer, &p.Executor, &p.TaxRate, &p.Account, &p.EndDate,
		&p.CreatedAt, &p.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("project not found")
	}
	if err != nil {
		return nil, err
	}

	calculateProject(p) // ← Вычисляем производные поля
	return p, nil
}

func (r *ProjectRepository) Update(userID, id int64, req *models.UpdateProjectRequest) (*models.Project, error) {
	current, err := r.GetByID(userID, id)
	if err != nil {
		return nil, err
	}

	// Применяем только непустые поля
	if req.Name != "" {
		current.Name = req.Name
	}
	if req.Priority != "" {
		current.Priority = req.Priority
	}
	if req.Stage != "" {
		current.Stage = req.Stage
	}
	if req.Start != "" {
		current.Start = req.Start
	}
	if req.Deadline != "" {
		current.Deadline = req.Deadline
	}
	if req.Cost != "" {
		current.Cost = req.Cost
	}
	if req.Paid != "" {
		current.Paid = req.Paid
	}
	if req.Remainder != "" {
		current.Remainder = req.Remainder
	}
	if req.OtherIncome != "" {
		current.OtherIncome = req.OtherIncome
	}
	if req.Revenue != "" {
		current.Revenue = req.Revenue
	}
	if req.Expense != "" {
		current.Expense = req.Expense
	}
	if req.Profit != "" {
		current.Profit = req.Profit
	}
	if req.Description != "" {
		current.Description = req.Description
	}
	if req.Customer != "" {
		current.Customer = req.Customer
	}
	if req.Executor != "" {
		current.Executor = req.Executor
	}
	if req.TaxRate != "" {
		current.TaxRate = req.TaxRate
	}
	if req.Account != "" {
		current.Account = req.Account
	}
	if req.EndDate != "" {
		current.EndDate = req.EndDate
	}
	current.UpdatedAt = time.Now()

	query := `
		UPDATE projects SET name=$1, priority=$2, stage=$3, start_date=$4, deadline=$5, cost=$6, paid=$7, remainder=$8,
		other_income=$9, revenue=$10, expense=$11, profit=$12, description=$13, customer=$14, executor=$15,
		tax_rate=$16, account=$17, end_date=$18, updated_at=$19
		WHERE id=$20 AND user_id=$21 RETURNING *`

	p := &models.Project{}
	err = r.db.QueryRow(context.Background(), query,
		current.Name, current.Priority, current.Stage, current.Start, current.Deadline,
		current.Cost, current.Paid, current.Remainder, current.OtherIncome, current.Revenue, current.Expense, current.Profit,
		current.Description, current.Customer, current.Executor, current.TaxRate, current.Account, current.EndDate,
		current.UpdatedAt, id, userID,
	).Scan(
		&p.ID, &p.UserID, &p.Name, &p.Priority, &p.Stage, &p.Start, &p.Deadline,
		&p.Cost, &p.Paid, &p.Remainder, &p.OtherIncome, &p.Revenue, &p.Expense, &p.Profit,
		&p.Description, &p.Customer, &p.Executor, &p.TaxRate, &p.Account, &p.EndDate,
		&p.CreatedAt, &p.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	calculateProject(p) // ← Вычисляем производные поля
	return p, nil
}

func (r *ProjectRepository) Delete(userID, id int64) error {
	result, err := r.db.Exec(context.Background(), `DELETE FROM projects WHERE id=$1 AND user_id=$2`, id, userID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("project not found")
	}
	return nil
}

func (r *ProjectRepository) DeleteMultiple(userID int64, ids []int64) error {
	_, err := r.db.Exec(context.Background(), `DELETE FROM projects WHERE user_id=$1 AND id = ANY($2)`, userID, ids)
	return err
}

// === Вспомогательные функции ===

// calculateProject вычисляет производные поля проекта
func calculateProject(p *models.Project) {
	// Вычисляем остаток: cost - paid
	cost := parseFloat(p.Cost)
	paid := parseFloat(p.Paid)
	remainder := cost - paid
	p.Remainder = formatMoney(remainder)

	// Вычисляем выручку: cost + other_income
	otherIncome := parseFloat(p.OtherIncome)
	revenue := cost + otherIncome
	p.Revenue = formatMoney(revenue)

	// Вычисляем прибыль: revenue - expense
	expense := parseFloat(p.Expense)
	profit := revenue - expense
	p.Profit = formatMoney(profit)

	// Вычисляем дни до дедлайна
	p.DaysLeft = daysUntil(p.EndDate)
}

// GetAllFiltered возвращает проекты с фильтрацией
func (r *ProjectRepository) GetAllFiltered(userID int64, filter models.ProjectFilter) ([]models.Project, error) {
	query := `SELECT * FROM projects WHERE user_id = $1`
	args := []interface{}{userID}
	argNum := 2

	// Фильтр по приоритету
	if filter.Priority != "" {
		query += fmt.Sprintf(" AND priority = $%d", argNum)
		args = append(args, filter.Priority)
		argNum++
	}

	// Фильтр по этапу
	if filter.Stage != "" {
		query += fmt.Sprintf(" AND stage = $%d", argNum)
		args = append(args, filter.Stage)
		argNum++
	}

	// Поиск по названию
	if filter.Search != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argNum)
		args = append(args, "%"+filter.Search+"%")
		argNum++
	}

	// Фильтр по дате начала
	if filter.StartDate != "" {
		query += fmt.Sprintf(" AND start_date >= $%d", argNum)
		args = append(args, filter.StartDate)
		argNum++
	}
	if filter.EndDate != "" {
		query += fmt.Sprintf(" AND start_date <= $%d", argNum)
		args = append(args, filter.EndDate)
		argNum++
	}

	// Сортировка
	switch filter.SortBy {
	case "priority":
		query += " ORDER BY priority"
	case "stage":
		query += " ORDER BY stage"
	case "deadline":
		query += " ORDER BY end_date"
	case "updated_at":
		query += " ORDER BY updated_at"
	default:
		query += " ORDER BY created_at"
	}

	if filter.Order == "asc" {
		query += " ASC"
	} else {
		query += " DESC"
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(
			&p.ID, &p.UserID, &p.Name, &p.Priority, &p.Stage, &p.Start, &p.Deadline,
			&p.Cost, &p.Paid, &p.Remainder, &p.OtherIncome, &p.Revenue, &p.Expense, &p.Profit,
			&p.Description, &p.Customer, &p.Executor, &p.TaxRate, &p.Account, &p.EndDate,
			&p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		calculateProject(&p)
		projects = append(projects, p)
	}

	if projects == nil {
		projects = []models.Project{}
	}
	return projects, nil
}

// parseFloat извлекает число из строки типа "35.000 р."
// func parseFloat(s string) float64 {
// 	s = strings.ReplaceAll(s, " р.", "")
// 	s = strings.ReplaceAll(s, ".", "")
// 	s = strings.ReplaceAll(s, ",", ".")
// 	s = strings.TrimSpace(s)

// 	if s == "" {
// 		return 0
// 	}

// 	val, err := strconv.ParseFloat(s, 64)
// 	if err != nil {
// 		return 0
// 	}
// 	return val
// }

// // formatMoney форматирует число в строку типа "35.000 р."
// func formatMoney(val float64) string {
// 	if val == 0 {
// 		return "0 р."
// 	}
// 	intPart := int(val)
// 	s := strconv.Itoa(intPart)
// 	if len(s) > 3 {
// 		s = s[:len(s)-3] + "." + s[len(s)-3:]
// 	}
// 	return s + " р."
// }

// // daysUntil возвращает количество дней до указанной даты
// func daysUntil(dateStr string) int {
// 	if dateStr == "" {
// 		return 0
// 	}

// 	var t time.Time
// 	var err error

// 	if strings.Contains(dateStr, "-") {
// 		t, err = time.Parse("2006-01-02", dateStr)
// 	} else {
// 		t, err = time.Parse("02.01.2006", dateStr)
// 	}

// 	if err != nil {
// 		return 0
// 	}

// 	days := int(time.Until(t).Hours() / 24)
// 	return days
// }
