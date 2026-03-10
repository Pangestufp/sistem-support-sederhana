package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type TicketDetailRepository interface {
	CreateTicketDetail(detail *entity.TicketDetail) error
	UpdateTicketDetail(detail *entity.TicketDetail) error
	DeleteTicketDetail(id int) error
	GetTicketDetailByID(id int) (*entity.TicketDetail, error)
	GetTicketDetailByTicketID(ticketID int) ([]entity.TicketDetail, error)
}

type ticketDetailRepository struct {
	db *gorm.DB
}

func NewTicketDetailRepository(db *gorm.DB) *ticketDetailRepository {
	return &ticketDetailRepository{db: db}
}

func (r *ticketDetailRepository) CreateTicketDetail(detail *entity.TicketDetail) error {
	return r.db.Create(detail).Error
}

func (r *ticketDetailRepository) UpdateTicketDetail(detail *entity.TicketDetail) error {
	result := r.db.Model(&entity.TicketDetail{}).
		Where("id = ?", detail.ID).
		Updates(map[string]interface{}{
			"user_id": detail.UserID,
			"review":  detail.Review,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketDetailRepository) DeleteTicketDetail(id int) error {
	result := r.db.Delete(&entity.TicketDetail{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketDetailRepository) GetTicketDetailByID(id int) (*entity.TicketDetail, error) {
	var detail entity.TicketDetail

	err := r.db.First(&detail, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &detail, nil
}

func (r *ticketDetailRepository) GetTicketDetailByTicketID(ticketID int) ([]entity.TicketDetail, error) {
	var details []entity.TicketDetail

	err := r.db.
		Where("ticket_id = ?", ticketID).
		Order("created_at ASC").
		Find(&details).Error

	if err != nil {
		return nil, err
	}

	return details, nil
}
