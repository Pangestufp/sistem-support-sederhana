package entity

import "time"

type TicketDetail struct {
	ID        int
	TicketID  int
	UserID    int
	Review    *string
	CreatedAt time.Time
}
