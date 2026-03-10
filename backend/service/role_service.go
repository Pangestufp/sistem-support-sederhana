package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RoleService interface {
	CreateRole(req *dto.CreateRoleRequest) error
	UpdateRole(roleID int, req *dto.UpdateRoleRequest) error
	DeleteRole(roleID int) error
	GetAllRole() ([]dto.RoleResponse, error)
	GetRoleByID(roleID int) (*dto.RoleResponse, error)
}

type roleService struct {
	repository repository.RoleRepository
	redis      *redis.Client
}

func NewRoleService(repository repository.RoleRepository, redis *redis.Client) *roleService {
	return &roleService{
		repository: repository,
		redis:      redis,
	}
}

func (s *roleService) CreateRole(req *dto.CreateRoleRequest) error {
	role := entity.Role{
		Name:        req.Name,
		Description: &req.Description,
		Status:      1,
	}
	return s.repository.CreateRole(&role)
}

func (s *roleService) UpdateRole(roleID int, req *dto.UpdateRoleRequest) error {
	role, err := s.repository.GetRoleByID(roleID)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	role.Name = req.Name
	role.Description = &req.Description

	return s.repository.UpdateRole(role)
}

func (s *roleService) DeleteRole(roleID int) error {
	return s.repository.DeleteRole(roleID)
}

func (s *roleService) GetAllRole() ([]dto.RoleResponse, error) {
	ctx := context.Background()
	cacheKey := "role:all"

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var responses []dto.RoleResponse
		json.Unmarshal([]byte(cached), &responses)
		return responses, nil
	}

	roles, err := s.repository.GetAllRole()
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.RoleResponse

	for _, role := range roles {
		desc := ""
		if role.Description != nil {
			desc = *role.Description
		}

		responses = append(responses, dto.RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			Description: desc,
		})
	}

	jsonData, _ := json.Marshal(responses)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return responses, nil
}

func (s *roleService) GetRoleByID(roleID int) (*dto.RoleResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("role:%d", roleID)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var response dto.RoleResponse
		json.Unmarshal([]byte(cached), &response)
		return &response, nil
	}

	role, err := s.repository.GetRoleByID(roleID)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	desc := ""
	if role.Description != nil {
		desc = *role.Description
	}

	response := dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: desc,
	}

	jsonData, _ := json.Marshal(response)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return &response, nil
}
