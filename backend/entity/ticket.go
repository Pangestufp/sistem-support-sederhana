package entity

import "time"

type Ticket struct {
	ID           int
	DocNo        string
	CreatedBy    int
	TicketTypeID int
	Description  *string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
