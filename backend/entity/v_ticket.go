package entity

import "time"

type VTicket struct {
	ID             int
	DocNo          string
	CreatedBy      int
	TicketTypeID   int
	TicketTypeName string
	Description    *string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time

	CreatorUsername string
	DepartmentID    int
	SuperiorID      *int
	LeadID          *int
}
