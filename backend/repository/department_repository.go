package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDeparment(department *entity.Department) error
	UpdateDeparment(department *entity.Department) error
	DeleteDeparment(id int) error
	GetDeparmentByID(id int) (*entity.Department, error)
	GetAllDepartment() ([]entity.Department, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *departmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) CreateDeparment(department *entity.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentRepository) UpdateDeparment(department *entity.Department) error {
	result := r.db.Model(&entity.Department{}).
		Where("id = ?", department.ID).
		Updates(map[string]interface{}{
			"name":        department.Name,
			"description": department.Description,
			"lead_id":     department.LeadID,
			"updated_at":  department.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *departmentRepository) DeleteDeparment(id int) error {
	result := r.db.Model(&entity.Department{}).
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

func (r *departmentRepository) GetDeparmentByID(id int) (*entity.Department, error) {
	var deparment entity.Department

	err := r.db.First(&deparment, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &deparment, nil
}

func (r *departmentRepository) GetAllDepartment() ([]entity.Department, error) {
	var deparments []entity.Department

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&deparments).Error
	if err != nil {
		return nil, err
	}

	return deparments, nil
}
