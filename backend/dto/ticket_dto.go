package dto

import (
	"mime/multipart"
	"time"
)

type TicketRequest struct {
	TicketTypeId int                     `form:"ticket_type_id"`
	Description  string                  `form:"description"`
	Pictures     []*multipart.FileHeader `form:"pictures"`
}

type TicketReviewRequest struct {
	Review   string                  `form:"review"`
	Pictures []*multipart.FileHeader `form:"pictures"`
}

type TicketResponse struct {
	ID             int    `json:"id"`
	DocNo          string `json:"doc_no"`
	CreatedBy      int    `json:"created_by"`
	TicketTypeID   int    `json:"ticket_type_id"`
	TicketTypeName string `json:"ticket_type_name"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type VTicketWorkflow struct {
	TicketID     int    `json:"ticket_id"`
	Step         string `json:"step"`
	AssignedUser string `json:"assigned_user"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Activity     string `json:"activity"`
	Action       string `json:"action"`
	ClosedAt     string `json:"closed_at"`
}

type UserTicketResponse struct {
	TicketID       int
	DocNo          string
	WorkflowPathID int
	ParallelKey    int
	AssignedUserID int
	ClosedAt       *time.Time
	Action         *string
	Activity       string
}
