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

type TicketTypeService interface {
	CreateTicketType(req *dto.CreateTicketTypeRequest) error
	UpdateTicketType(id int, req *dto.UpdateTicketTypeRequest) error
	DeleteTicketType(id int) error
	GetAllTicketType() ([]dto.TicketTypeResponse, error)
	GetTicketTypeByID(id int) (*dto.TicketTypeResponse, error)
}

type ticketTypeService struct {
	repository repository.TicketTypeRepository
	redis      *redis.Client
}

func NewTicketTypeService(repository repository.TicketTypeRepository, redis *redis.Client) *ticketTypeService {
	return &ticketTypeService{
		repository: repository,
		redis:      redis,
	}
}

func (s *ticketTypeService) CreateTicketType(req *dto.CreateTicketTypeRequest) error {

	tt := entity.TicketType{
		Name:        req.Name,
		Description: &req.Description,
		Status:      1,
	}

	return s.repository.CreateTicketType(&tt)
}

func (s *ticketTypeService) UpdateTicketType(id int, req *dto.UpdateTicketTypeRequest) error {

	tt, err := s.repository.GetTicketTypeByID(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	tt.Name = req.Name
	tt.Description = &req.Description

	return s.repository.UpdateTicketType(tt)
}

func (s *ticketTypeService) DeleteTicketType(id int) error {
	return s.repository.DeleteTicketType(id)
}

func (s *ticketTypeService) GetAllTicketType() ([]dto.TicketTypeResponse, error) {
	ctx := context.Background()
	cacheKey := "ticket_type:all"

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var responses []dto.TicketTypeResponse
		json.Unmarshal([]byte(cached), &responses)
		return responses, nil
	}

	list, err := s.repository.GetAllTicketType()
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.TicketTypeResponse

	for _, tt := range list {

		desc := ""
		if tt.Description != nil {
			desc = *tt.Description
		}

		responses = append(responses, dto.TicketTypeResponse{
			ID:          tt.ID,
			Name:        tt.Name,
			Description: desc,
		})
	}

	jsonData, _ := json.Marshal(responses)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return responses, nil
}

func (s *ticketTypeService) GetTicketTypeByID(id int) (*dto.TicketTypeResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("ticket_type:%d", id)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var response dto.TicketTypeResponse
		json.Unmarshal([]byte(cached), &response)
		return &response, nil
	}

	tt, err := s.repository.GetTicketTypeByID(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	desc := ""
	if tt.Description != nil {
		desc = *tt.Description
	}

	response := dto.TicketTypeResponse{
		ID:          tt.ID,
		Name:        tt.Name,
		Description: desc,
	}

	jsonData, _ := json.Marshal(response)
	s.redis.Set(ctx, cacheKey, jsonData, 5*time.Minute)

	return &response, nil
}
