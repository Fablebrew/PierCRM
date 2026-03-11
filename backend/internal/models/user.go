package models

import "time"

type User struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	BusinessForm   string    `json:"business_form"`
	NalogSystem    string    `json:"nalog_system"`
	Employees      string    `json:"employees"`
	BusinessSphere string    `json:"business_sphere"`
	Password       string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}
