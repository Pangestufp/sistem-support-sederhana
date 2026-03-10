package repository

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TicketWorkflowRepository interface {
	CreateTicketWorkflow(workflow *entity.TicketWorkflow) error
	CreateTicketWorkflows(workflows []entity.TicketWorkflow) error
	CloseTicketWorkflow(ticketWorkflow entity.TicketWorkflow) error
	GetTicketWorkflowByID(id int) (*entity.TicketWorkflow, error)
	GetTicketWorkflowByTicketID(ticketID int) ([]entity.TicketWorkflow, error)
	GetCurrentTicketWorkflowByTicketID(ticketID, userID int) (*entity.TicketWorkflow, error)
	GetCurrentWorkflow(userID int) ([]entity.VUserTicket, error)
	ReopenLastWorkflow(ticketID int) error
	EnsureClosedIfWorkflowFinished(ticketID int) error
	GetTicketWorkflowV(ticketID int) ([]dto.VTicketWorkflow, error)
}

type ticketWorkflowRepository struct {
	db *gorm.DB
}

func NewTicketWorkflowRepository(db *gorm.DB) *ticketWorkflowRepository {
	return &ticketWorkflowRepository{db: db}
}

func (r *ticketWorkflowRepository) CreateTicketWorkflow(workflow *entity.TicketWorkflow) error {
	return r.db.Create(workflow).Error
}

func (r *ticketWorkflowRepository) CreateTicketWorkflows(workflows []entity.TicketWorkflow) error {
	if len(workflows) == 0 {
		return nil
	}

	return r.db.Create(&workflows).Error
}

func (r *ticketWorkflowRepository) CloseTicketWorkflow(ticketWorkflow entity.TicketWorkflow) error {
	result := r.db.
		Model(&entity.TicketWorkflow{}).
		Where(
			"ticket_id = ? AND workflow_path_id = ? AND parallel_key = ? AND closed_at IS NULL",
			ticketWorkflow.TicketID,
			ticketWorkflow.WorkflowPathID,
			ticketWorkflow.ParallelKey,
		).
		Updates(map[string]interface{}{
			"closed_at": ticketWorkflow.ClosedAt,
			"action":    ticketWorkflow.Action,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketWorkflowRepository) GetTicketWorkflowByID(id int) (*entity.TicketWorkflow, error) {
	var workflow entity.TicketWorkflow

	err := r.db.First(&workflow, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &workflow, nil
}

func (r *ticketWorkflowRepository) GetTicketWorkflowByTicketID(ticketID int) ([]entity.TicketWorkflow, error) {
	var workflows []entity.TicketWorkflow

	err := r.db.
		Where("ticket_id = ?", ticketID).
		Order("workflow_path_id ASC, parallel_key ASC").
		Find(&workflows).Error

	if err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *ticketWorkflowRepository) GetCurrentTicketWorkflowByTicketID(ticketID, userID int) (*entity.TicketWorkflow, error) {
	var workflows []entity.TicketWorkflow

	err := r.db.
		Where("ticket_id = ? AND closed_at IS NULL AND NOT EXISTS (SELECT 1 FROM ticket_workflows tw WHERE tw.ticket_id = ticket_workflows.ticket_id AND tw.action = 'reject')", ticketID).
		Order("workflow_path_id ASC, parallel_key ASC").
		Find(&workflows).Error

	if err != nil {
		return nil, err
	}

	if len(workflows) == 0 {
		return nil, &errorhandler.NotFoundError{Message: "Data not found"}
	}

	first := workflows[0]

	for _, wf := range workflows {
		if wf.WorkflowPathID != first.WorkflowPathID || wf.ParallelKey != first.ParallelKey {
			break
		}
		if wf.AssignedUserID == userID {
			return &wf, nil
		}
	}

	return &first, nil
}

func (r *ticketWorkflowRepository) GetCurrentWorkflow(userID int) ([]entity.VUserTicket, error) {
	var workflows []entity.VUserTicket

	sql := `select t.*, th.doc_no from (
		SELECT *
		FROM (
			SELECT
				wf.*,
				RANK() OVER (
					PARTITION BY wf.ticket_id
					ORDER BY wf.workflow_path_id, wf.parallel_key DESC
				) AS urutan
			FROM ticket_workflows wf
			WHERE wf.closed_at IS NULL
			AND NOT EXISTS (
					SELECT 1
					FROM ticket_workflows tw
					WHERE tw.ticket_id = wf.ticket_id
					  AND tw.action = 'reject'
				)
		) h
		WHERE h.urutan = 1 AND h.assigned_user_id = ? ) t join tickets th on t.ticket_id = th.id
	`
	err := r.db.Raw(sql, userID).Scan(&workflows).Error
	if err != nil {
		return nil, err
	}

	return workflows, nil
}

func (r *ticketWorkflowRepository) ReopenLastWorkflow(ticketID int) error {

	var wf entity.TicketWorkflow
	err := r.db.
		Table("ticket_workflows").
		Where(`
            ticket_id = ?
            AND closed_at IS NOT NULL
            AND NOT EXISTS (
                SELECT 1
                FROM ticket_workflows tw
                WHERE tw.ticket_id = ticket_workflows.ticket_id
                AND tw.action = 'reject'
            )
        `, ticketID).
		Order("workflow_path_id DESC, parallel_key DESC").
		Limit(1).
		Take(&wf).Error

	if err != nil {
		return err
	}

	err = r.db.
		Model(&entity.TicketWorkflow{}).
		Where(
			"ticket_id = ? AND workflow_path_id = ? AND parallel_key = ?",
			ticketID,
			wf.WorkflowPathID,
			wf.ParallelKey,
		).
		Updates(map[string]interface{}{
			"closed_at": nil,
			"action":    nil,
		}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *ticketWorkflowRepository) EnsureClosedIfWorkflowFinished(ticketID int) error {
	var count int64

	err := r.db.
		Model(&entity.TicketWorkflow{}).
		Where("ticket_id = ? AND closed_at IS NULL", ticketID).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	err = r.db.
		Model(&entity.Ticket{}).
		Where("id = ?", ticketID).
		Updates(map[string]interface{}{
			"status":     "closed",
			"updated_at": time.Now(),
		}).Error

	return err
}

func (r *ticketWorkflowRepository) GetTicketWorkflowV(ticketID int) ([]dto.VTicketWorkflow, error) {
	var workflows []dto.VTicketWorkflow

	err := r.db.
		Where("ticket_id = ?", ticketID).
		Find(&workflows).Error

	if err != nil {
		return nil, err
	}

	return workflows, nil
}
