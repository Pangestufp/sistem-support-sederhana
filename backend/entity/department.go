package entity

import "time"

type Department struct {
	ID          int
	Name        string
	Description *string
	LeadID      *int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      int
}
