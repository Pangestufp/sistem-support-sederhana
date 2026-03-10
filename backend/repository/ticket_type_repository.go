package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type TicketTypeRepository interface {
	CreateTicketType(ticketType *entity.TicketType) error
	UpdateTicketType(ticketType *entity.TicketType) error
	DeleteTicketType(id int) error
	GetTicketTypeByID(id int) (*entity.TicketType, error)
	GetAllTicketType() ([]entity.TicketType, error)
}

type ticketTypeRepository struct {
	db *gorm.DB
}

func NewTicketTypeRepository(db *gorm.DB) *ticketTypeRepository {
	return &ticketTypeRepository{
		db: db,
	}
}

func (r *ticketTypeRepository) CreateTicketType(ticketType *entity.TicketType) error {
	return r.db.Create(ticketType).Error
}

func (r *ticketTypeRepository) UpdateTicketType(ticketType *entity.TicketType) error {
	result := r.db.Model(&entity.TicketType{}).
		Where("id = ?", ticketType.ID).
		Updates(map[string]interface{}{
			"name":        ticketType.Name,
			"description": ticketType.Description,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketTypeRepository) DeleteTicketType(id int) error {
	result := r.db.Model(&entity.TicketType{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *ticketTypeRepository) GetTicketTypeByID(id int) (*entity.TicketType, error) {
	var ticketType entity.TicketType

	err := r.db.First(&ticketType, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &ticketType, nil
}

func (r *ticketTypeRepository) GetAllTicketType() ([]entity.TicketType, error) {
	var ticketTypes []entity.TicketType

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&ticketTypes).Error
	if err != nil {
		return nil, err
	}

	return ticketTypes, nil
}
