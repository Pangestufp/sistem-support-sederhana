package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	UpdateUser(User *entity.User) error
	GetUserByID(id int) (*entity.User, error)
	GetAllUser() ([]entity.User, error)
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) UpdateUser(user *entity.User) error {
	result := r.db.Model(&entity.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":          user.Name,
			"email":         user.Email,
			"username":      user.Username,
			"department_id": user.DepartmentID,
			"superior_id":   user.SuperiorID,
			"status":        user.Status,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *userRepository) DeleteUser(id int) error {
	result := r.db.Model(&entity.User{}).
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

func (r *userRepository) GetUserByID(id int) (*entity.User, error) {
	var User entity.User

	err := r.db.First(&User, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &User, nil
}

func (r *userRepository) GetAllUser() ([]entity.User, error) {
	var Users []entity.User

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&Users).Error
	if err != nil {
		return nil, err
	}

	return Users, nil
}
