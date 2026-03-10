package service

import (
	"TicketManagement/dto"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/repository"
	"time"
)

type UserService interface {
	UpdateUser(userID int, req *dto.UpdateUserRequest) error
	UpdatePassword(userID int, req *dto.UpdatePasswordRequest) error
	DeleteUser(userID int) error
	GetAllUser() ([]dto.UserResponse, error)
	GetUserByID(userID int) (*dto.UserResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) UpdateUser(userID int, req *dto.UpdateUserRequest) error {

	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "User Not Found"}
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.DepartmentID = req.DepartmentID
	user.SuperiorID = req.SuperiorID
	user.Status = req.Status
	user.UpdatedAt = time.Now()

	return s.repository.UpdateUser(user)
}

func (s *userService) UpdatePassword(userID int, req *dto.UpdatePasswordRequest) error {

	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "User Not Found"}
	}

	if err := helper.VerifyPassword(user.Password, req.OldPassword); err != nil {
		return &errorhandler.BadRequestError{Message: "Old password is incorrect"}
	}

	if req.NewPassword != req.NewPasswordConfirm {
		return &errorhandler.BadRequestError{Message: "Password confirmation does not match"}
	}

	newHashedPassword, err := helper.HashPassword(req.NewPassword)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user.Password = newHashedPassword
	user.UpdatedAt = time.Now()

	return s.repository.UpdateUser(user)
}

func (s *userService) DeleteUser(userID int) error {
	return s.repository.DeleteUser(userID)
}

func (s *userService) GetAllUser() ([]dto.UserResponse, error) {

	users, err := s.repository.GetAllUser()
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.UserResponse

	for _, user := range users {

		responses = append(responses, dto.UserResponse{
			ID:           user.ID,
			Name:         user.Name,
			Username:     user.Username,
			Email:        user.Email,
			DepartmentID: user.DepartmentID,
			SuperiorID:   user.SuperiorID,
			Status:       user.Status,
			CreatedAt:    helper.FormatTimeRFC3339(user.CreatedAt),
			UpdatedAt:    helper.FormatTimeRFC3339(user.UpdatedAt),
		})
	}

	return responses, nil
}

func (s *userService) GetUserByID(userID int) (*dto.UserResponse, error) {

	user, err := s.repository.GetUserByID(userID)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "User Not Found"}
	}

	response := dto.UserResponse{
		ID:           user.ID,
		Name:         user.Name,
		Username:     user.Username,
		Email:        user.Email,
		DepartmentID: user.DepartmentID,
		SuperiorID:   user.SuperiorID,
		Status:       user.Status,
		CreatedAt:    helper.FormatTimeRFC3339(user.CreatedAt),
		UpdatedAt:    helper.FormatTimeRFC3339(user.UpdatedAt),
	}

	return &response, nil
}
