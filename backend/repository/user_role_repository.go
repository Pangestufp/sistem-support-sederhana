package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"

	"gorm.io/gorm"
)

type UserRoleRepository interface {
	CreateUserRole(userRole *entity.UserRole) error
	DeleteUserRole(userID, roleID int) error
	GetAllUserRole(userID int) ([]entity.UserRole, error)
	GetUserByRoleName(role string) ([]entity.VUserRole, error)
}

type userRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *userRoleRepository {
	return &userRoleRepository{
		db: db,
	}
}

func (r *userRoleRepository) CreateUserRole(userRole *entity.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *userRoleRepository) DeleteUserRole(userID, roleID int) error {
	result := r.db.
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&entity.UserRole{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil

}

func (r *userRoleRepository) GetAllUserRole(userID int) ([]entity.UserRole, error) {
	var userRole []entity.UserRole

	err := r.db.Where("user_id = ?", userID).Find(&userRole).Error
	if err != nil {
		return nil, err
	}

	return userRole, nil
}

func (r *userRoleRepository) GetUserByRoleName(role string) ([]entity.VUserRole, error) {
	var VUserRoles []entity.VUserRole

	err := r.db.Where("Name = ?", role).Find(&VUserRoles).Error

	if err != nil {
		return nil, err
	}

	return VUserRoles, err

}
