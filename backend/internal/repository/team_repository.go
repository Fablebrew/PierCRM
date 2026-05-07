package repository

import (
	"context"
	"finances/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeamRepository struct {
	db *pgxpool.Pool
}

func NewTeamRepository(db *pgxpool.Pool) *TeamRepository {
	return &TeamRepository{db: db}
}

// Create - создание нового члена команды
func (r *TeamRepository) Create(userID int64, req *models.CreateTeamMemberRequest) (*models.TeamMember, error) {
	query := `
		INSERT INTO team_members (user_id, name, email, phone, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, user_id, name, email, phone, role, created_at, updated_at
	`

	now := time.Now()
	member := &models.TeamMember{}

	err := r.db.QueryRow(
		context.Background(),
		query,
		userID,
		req.Name,
		req.Email,
		req.Phone,
		req.Role,
		now,
		now,
	).Scan(
		&member.ID,
		&member.UserID,
		&member.Name,
		&member.Email,
		&member.Phone,
		&member.Role,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err != nil {
		// Проверяем на дубликат email (код ошибки уникальности в PostgreSQL)
		if strings.Contains(err.Error(), "duplicate key") ||
			strings.Contains(err.Error(), "unique_user_email") {
			return nil, fmt.Errorf("team member with this email already exists")
		}
		return nil, fmt.Errorf("failed to create team member: %w", err)
	}

	return member, nil
}

// GetByID - получение члена команды по ID
func (r *TeamRepository) GetByID(userID, memberID int64) (*models.TeamMember, error) {
	query := `
		SELECT id, user_id, name, email, phone, role, created_at, updated_at
		FROM team_members
		WHERE id = $1 AND user_id = $2
	`

	member := &models.TeamMember{}
	err := r.db.QueryRow(
		context.Background(),
		query,
		memberID,
		userID,
	).Scan(
		&member.ID,
		&member.UserID,
		&member.Name,
		&member.Email,
		&member.Phone,
		&member.Role,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("team member not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get team member: %w", err)
	}

	return member, nil
}

// GetAll - получение всех членов команды пользователя
func (r *TeamRepository) GetAll(userID int64) ([]models.TeamMember, error) {
	query := `
		SELECT id, user_id, name, email, phone, role, created_at, updated_at
		FROM team_members
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get team members: %w", err)
	}
	defer rows.Close()

	var members []models.TeamMember
	for rows.Next() {
		var member models.TeamMember
		err := rows.Scan(
			&member.ID,
			&member.UserID,
			&member.Name,
			&member.Email,
			&member.Phone,
			&member.Role,
			&member.CreatedAt,
			&member.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan team member: %w", err)
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating team members: %w", err)
	}

	// Возвращаем пустой массив вместо nil для JSON
	if members == nil {
		members = []models.TeamMember{}
	}

	return members, nil
}

// Update - обновление члена команды
func (r *TeamRepository) Update(userID, memberID int64, req *models.UpdateTeamMemberRequest) (*models.TeamMember, error) {
	// Сначала получаем текущие данные
	current, err := r.GetByID(userID, memberID)
	if err != nil {
		return nil, err
	}

	// Применяем новые значения, если они предоставлены
	if req.Name != "" {
		current.Name = req.Name
	}
	if req.Email != "" {
		current.Email = req.Email
	}
	if req.Phone != "" {
		current.Phone = req.Phone
	}
	if req.Role != "" {
		current.Role = req.Role
	}
	current.UpdatedAt = time.Now()

	query := `
		UPDATE team_members 
		SET name = $1, email = $2, phone = $3, role = $4, updated_at = $5
		WHERE id = $6 AND user_id = $7
		RETURNING id, user_id, name, email, phone, role, created_at, updated_at
	`

	member := &models.TeamMember{}
	err = r.db.QueryRow(
		context.Background(),
		query,
		current.Name,
		current.Email,
		current.Phone,
		current.Role,
		current.UpdatedAt,
		memberID,
		userID,
	).Scan(
		&member.ID,
		&member.UserID,
		&member.Name,
		&member.Email,
		&member.Phone,
		&member.Role,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, fmt.Errorf("team member with this email already exists")
		}
		return nil, fmt.Errorf("failed to update team member: %w", err)
	}

	return member, nil
}

// Delete - удаление члена команды
func (r *TeamRepository) Delete(userID, memberID int64) error {
	query := `DELETE FROM team_members WHERE id = $1 AND user_id = $2`

	ct, err := r.db.Exec(context.Background(), query, memberID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete team member: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("team member not found or not authorized")
	}

	return nil
}

// DeleteMultiple - удаление нескольких членов команды
func (r *TeamRepository) DeleteMultiple(userID int64, memberIDs []int64) error {
	if len(memberIDs) == 0 {
		return fmt.Errorf("no member IDs provided")
	}

	query := `DELETE FROM team_members WHERE user_id = $1 AND id = ANY($2)`

	_, err := r.db.Exec(context.Background(), query, userID, memberIDs)
	if err != nil {
		return fmt.Errorf("failed to delete team members: %w", err)
	}

	return nil
}
