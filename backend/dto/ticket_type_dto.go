package dto

type CreateTicketTypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTicketTypeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TicketTypeResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
