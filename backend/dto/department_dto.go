package dto

type CreateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LeadID      *int   `json:"lead_id"`
}

type UpdateDepartmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LeadID      *int   `json:"lead_id"`
}

type DepartmentResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	LeadID      *int   `json:"lead_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
