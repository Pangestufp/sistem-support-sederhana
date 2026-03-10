package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/repository"
)

type WorkflowService interface {
	CreateWorkflow(req *dto.CreateWorkflowRequest) error
	UpdateWorkflow(id int, req *dto.UpdateWorkflowRequest) error
	DeleteWorkflow(id int) error
	GetAllWorkflow() ([]dto.WorkflowResponse, error)
	GetWorkflowByID(id int) (*dto.WorkflowResponse, error)
}

type workflowService struct {
	repository repository.WorkflowRepository
}

func NewWorkflowService(repository repository.WorkflowRepository) *workflowService {
	return &workflowService{
		repository: repository,
	}
}

func (s *workflowService) CreateWorkflow(req *dto.CreateWorkflowRequest) error {
	workflow := entity.Workflow{
		Name:        req.Name,
		Description: &req.Description,
	}

	s.repository.CreateWorkflow(&workflow)
	return nil
}

func (s *workflowService) UpdateWorkflow(id int, req *dto.UpdateWorkflowRequest) error {
	workflow, err := s.repository.GetWorkflowByID(id)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	workflow.Name = req.Name
	workflow.Description = &req.Description

	s.repository.UpdateWorkflow(workflow)
	return nil
}

func (s *workflowService) DeleteWorkflow(id int) error {
	return s.repository.DeleteWorkflow(id)
}

func (s *workflowService) GetAllWorkflow() ([]dto.WorkflowResponse, error) {
	workflows, err := s.repository.GetAllWorkflow()
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.WorkflowResponse

	for _, wf := range workflows {
		desc := ""
		if wf.Description != nil {
			desc = *wf.Description
		}

		responses = append(responses, dto.WorkflowResponse{
			ID:          wf.ID,
			Name:        wf.Name,
			Description: desc,
		})
	}

	return responses, nil
}

func (s *workflowService) GetWorkflowByID(id int) (*dto.WorkflowResponse, error) {
	wf, err := s.repository.GetWorkflowByID(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	desc := ""
	if wf.Description != nil {
		desc = *wf.Description
	}

	response := dto.WorkflowResponse{
		ID:          wf.ID,
		Name:        wf.Name,
		Description: desc,
	}

	return &response, nil
}
