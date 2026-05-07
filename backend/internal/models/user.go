package models

import "time"

type User struct {
	ID             int64     `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Email          string    `json:"email" db:"email"`
	PasswordHash   string    `json:"password_hash,omitempty" db:"password_hash"`
	BusinessForm   string    `json:"business_form,omitempty" db:"business_form"`
	BusinessSphere string    `json:"business_sphere,omitempty" db:"business_sphere"`
	Employees      string    `json:"employees,omitempty" db:"employees"`
	NalogSystem    string    `json:"nalog_system,omitempty" db:"nalog_system"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

type RegisterRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=8"`
	BusinessForm   string `json:"business_form"`
	BusinessSphere string `json:"business_sphere"`
	Employees      string `json:"employees"`
	NalogSystem    string `json:"nalog_system"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	BusinessForm   string    `json:"business_form,omitempty"`
	BusinessSphere string    `json:"business_sphere,omitempty"`
	Employees      string    `json:"employees,omitempty"`
	NalogSystem    string    `json:"nalog_system,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
