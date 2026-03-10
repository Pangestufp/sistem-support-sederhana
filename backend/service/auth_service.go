package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/repository"
	"time"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repositoryA repository.AuthRepository
	repositoryR repository.UserRoleRepository
}

func NewAuthService(repositoryA repository.AuthRepository, repositoryR repository.UserRoleRepository) *authService {
	return &authService{
		repositoryA: repositoryA,
		repositoryR: repositoryR,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repositoryA.EmailExist(req.Email); emailExist {
		return &errorhandler.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password)

	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:         req.Name,
		Email:        req.Email,
		Password:     passwordHash,
		Username:     req.Username,
		DepartmentID: req.DepartmentID,
		SuperiorID:   req.SuperiorID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Status:       1,
	}

	if err := s.repositoryA.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repositoryA.GetUserByEmail(req.Email)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "wrong email or password"}
	}

	userRole, err := s.repositoryR.GetAllUserRole(user.ID)

	if err != nil {
		userRole = []entity.UserRole{}
	}

	token, err := helper.GenerateToken(user, userRole)

	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	res := dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &res, err
}
