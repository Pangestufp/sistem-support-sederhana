package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type TicketRepository interface {
	CreateTicket(ticket *entity.Ticket) error
	UpdateTicket(ticket *entity.Ticket) error
	GetTicketByID(id int) (*entity.Ticket, error)
	GetAllTicket() ([]entity.Ticket, error)
	IsConditionMet(id int, readColumn string, exeCondition string) (bool, error)
	GetVTicketID(id int) (*entity.VTicket, error)
	GenerateDocNo() (string, error)
	GetTicketWithPagination(ticketID int) ([]entity.Ticket, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (r *ticketRepository) CreateTicket(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) UpdateTicket(ticket *entity.Ticket) error {
	result := r.db.Model(&entity.Ticket{}).
		Where("id = ?", ticket.ID).
		Updates(map[string]interface{}{
			"status":     ticket.Status,
			"updated_at": time.Now(),
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketRepository) GetTicketByID(id int) (*entity.Ticket, error) {
	var ticket entity.Ticket

	err := r.db.First(&ticket, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &ticket, nil
}

func (r *ticketRepository) GetAllTicket() ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&tickets).Error
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *ticketRepository) IsConditionMet(id int, readColumn string, exeCondition string) (bool, error) {

	query := fmt.Sprintf(`
		SELECT 1
		FROM v_tickets
		WHERE id = ?
		  AND %s %s
		LIMIT 1
	`, readColumn, exeCondition)

	log.Printf("Executing query: %s | id=%d", query, id)

	var dummy int
	tx := r.db.Raw(query, id).Scan(&dummy)

	if tx.Error != nil {
		return false, tx.Error
	}

	if tx.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (r *ticketRepository) GetVTicketID(id int) (*entity.VTicket, error) {
	var ticket entity.VTicket

	err := r.db.First(&ticket, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &ticket, nil
}

func (r *ticketRepository) GenerateDocNo() (string, error) {
	var docNo string

	err := r.db.
		Raw("CALL generate_doc_no()").
		Scan(&docNo).
		Error

	if err != nil {
		return "", err
	}

	return docNo, nil
}

func (r *ticketRepository) GetTicketWithPagination(lastID int) ([]entity.Ticket, error) {
	var tickets []entity.Ticket

	err := r.db.Where("id > ?", lastID).
		Order("id ASC").
		Limit(10).Find(&tickets).Error
	if err != nil {
		return nil, err
	}

	return tickets, nil
}
