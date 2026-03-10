package dto

type TicketDetailResponse struct {
	ID        int     `json:"id"`
	TicketID  int     `json:"ticket_id"`
	UserID    int     `json:"user_id"`
	Review    *string `json:"review"`
	CreatedAt string  `json:"created_at"`
}
