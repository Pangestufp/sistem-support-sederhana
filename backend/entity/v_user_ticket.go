package entity

import "time"

type VUserTicket struct {
	TicketID       int
	DocNo          string
	WorkflowPathID int
	ParallelKey    int
	AssignedUserID int
	ClosedAt       *time.Time
	Action         *string
	Activity       string
}
