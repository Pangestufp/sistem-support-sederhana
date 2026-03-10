package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type TicketAttachmentRepository interface {
	CreateTicketAttachment(attachment *entity.TicketAttachment) error
	DeleteTicketAttachment(id int) error
	GetTicketAttachmentByID(id int) (*entity.TicketAttachment, error)
	GetTicketAttachmentByTicketID(ticketID int) ([]entity.TicketAttachment, error)
}

type ticketAttachmentRepository struct {
	db *gorm.DB
}

func NewTicketAttachmentRepository(db *gorm.DB) *ticketAttachmentRepository {
	return &ticketAttachmentRepository{db: db}
}

func (r *ticketAttachmentRepository) CreateTicketAttachment(attachment *entity.TicketAttachment) error {
	return r.db.Create(attachment).Error
}

func (r *ticketAttachmentRepository) DeleteTicketAttachment(id int) error {
	result := r.db.Delete(&entity.TicketAttachment{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketAttachmentRepository) GetTicketAttachmentByID(id int) (*entity.TicketAttachment, error) {
	var attachment entity.TicketAttachment

	err := r.db.First(&attachment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &attachment, nil
}

func (r *ticketAttachmentRepository) GetTicketAttachmentByTicketID(ticketID int) ([]entity.TicketAttachment, error) {
	var attachments []entity.TicketAttachment

	err := r.db.
		Where("ticket_id = ?", ticketID).
		Order("created_at ASC").
		Find(&attachments).Error

	if err != nil {
		return nil, err
	}

	return attachments, nil
}
