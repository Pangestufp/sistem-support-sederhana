package entity

import "time"

type TicketAttachment struct {
	ID        int
	TicketID  int
	FilePath  string
	Note      *string
	CreatedAt time.Time
}
