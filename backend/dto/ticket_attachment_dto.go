package dto

type TicketAttachmentResponse struct {
	ID        int     `json:"id"`
	TicketID  int     `json:"ticket_id"`
	FilePath  string  `json:"file_path"`
	Note      *string `json:"note"`
	CreatedAt string  `json:"created_at"`
}
