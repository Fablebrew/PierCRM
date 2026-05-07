package models

import "time"

// TeamMember - модель члена команды
type TeamMember struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Phone     string    `json:"phone" db:"phone"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateTeamMemberRequest - запрос на создание
type CreateTeamMemberRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=100"`
	Email string `json:"email" binding:"required,email"`
	Phone string `json:"phone"`
	Role  string `json:"role" binding:"max=100"`
}

// UpdateTeamMemberRequest - запрос на обновление
type UpdateTeamMemberRequest struct {
	Name  string `json:"name" binding:"omitempty,min=2,max=100"`
	Email string `json:"email" binding:"omitempty,email"`
	Phone string `json:"phone"`
	Role  string `json:"role" binding:"omitempty,max=100"`
}

// BulkDeleteRequest - запрос на массовое удаление
type BulkDeleteRequest struct {
	IDs []int64 `json:"ids" binding:"required,min=1"`
}
