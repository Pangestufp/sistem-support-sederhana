package entity

import "time"

type User struct {
	ID           int
	Username     string
	Password     string
	Email        string
	Name         string
	DepartmentID *int
	SuperiorID   *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Status       int
}
