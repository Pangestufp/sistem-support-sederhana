package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/repository"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type DepartmentService interface {
	CreateDepartment(req *dto.CreateDepartmentRequest) error
	UpdateDepartment(departmentID int, req *dto.UpdateDepartmentRequest) error
	DeleteDepartment(departmentID int) error
	GetallDepartment() ([]dto.DepartmentResponse, error)
	GetDepartmentByID(departmentID int) (*dto.DepartmentResponse, error)
}

type departmentService struct {
	repository repository.DepartmentRepository
	redis      *redis.Client
}

func NewDeparmentService(repository repository.DepartmentRepository, redis *redis.Client) *departmentService {
	return &departmentService{
		repository: repository,
		redis:      redis,
	}
}

func (s *departmentService) CreateDepartment(req *dto.CreateDepartmentRequest) error {
	department := entity.Department{
		Name:        req.Name,
		Description: &req.Description,
		LeadID:      req.LeadID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      1,
	}
	return s.repository.CreateDeparment(&department)
}

func (s *departmentService) UpdateDepartment(departmentID int, req *dto.UpdateDepartmentRequest) error {
	department, err := s.repository.GetDeparmentByID(departmentID)

	if err != nil {
		return &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	department.Name = req.Name
	department.Description = &req.Description
	department.LeadID = req.LeadID
	department.UpdatedAt = time.Now()

	return s.repository.UpdateDeparment(department)
}

func (s *departmentService) DeleteDepartment(departmentID int) error {
	return s.repository.DeleteDeparment(departmentID)
}

func (s *departmentService) GetallDepartment() ([]dto.DepartmentResponse, error) {

	ctx := context.Background()
	cacheKey := "department:all"

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var responses []dto.DepartmentResponse
		json.Unmarshal([]byte(cached), &responses)
		return responses, nil
	}

	departments, err := s.repository.GetAllDepartment()

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.DepartmentResponse

	for _, dept := range departments {
		description := ""
		if dept.Description != nil {
			description = *dept.Description
		}
		responses = append(
			responses, dto.DepartmentResponse{
				ID:          dept.ID,
				Name:        dept.Name,
				Description: description,
				LeadID:      dept.LeadID,
				CreatedAt:   helper.FormatTimeRFC3339(dept.CreatedAt),
				UpdatedAt:   helper.FormatTimeRFC3339(dept.UpdatedAt),
			})
	}

	jsonData, _ := json.Marshal(responses)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return responses, nil
}

func (s *departmentService) GetDepartmentByID(departmentID int) (*dto.DepartmentResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("department:%d", departmentID)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var response dto.DepartmentResponse
		json.Unmarshal([]byte(cached), &response)
		return &response, nil
	}

	department, err := s.repository.GetDeparmentByID(departmentID)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	description := ""

	if department.Description != nil {
		description = *department.Description
	}

	response := dto.DepartmentResponse{
		ID:          department.ID,
		Name:        department.Name,
		Description: description,
		LeadID:      department.LeadID,
		CreatedAt:   helper.FormatTimeRFC3339(department.CreatedAt),
		UpdatedAt:   helper.FormatTimeRFC3339(department.UpdatedAt),
	}

	jsonData, _ := json.Marshal(response)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return &response, nil
}
