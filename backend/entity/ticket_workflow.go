package entity

import "time"

type TicketWorkflow struct {
	TicketID       int
	WorkflowPathID int
	ParallelKey    int
	AssignedUserID int
	ClosedAt       *time.Time
	Action         *string
	Activity       string
}
