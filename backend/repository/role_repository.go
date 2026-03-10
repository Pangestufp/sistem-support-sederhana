package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *entity.Role) error
	UpdateRole(role *entity.Role) error
	GetRoleByID(id int) (*entity.Role, error)
	GetAllRole() ([]entity.Role, error)
	DeleteRole(id int) error
}

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *roleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) CreateRole(role *entity.Role) error {
	return r.db.Create(role).Error
}

func (r *roleRepository) UpdateRole(role *entity.Role) error {
	result := r.db.Model(&entity.Role{}).
		Where("id = ?", role.ID).
		Updates(map[string]interface{}{
			"name":        role.Name,
			"description": role.Description,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *roleRepository) DeleteRole(id int) error {
	result := r.db.Model(&entity.Role{}).
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

func (r *roleRepository) GetRoleByID(id int) (*entity.Role, error) {
	var role entity.Role

	err := r.db.First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) GetAllRole() ([]entity.Role, error) {
	var roles []entity.Role

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
