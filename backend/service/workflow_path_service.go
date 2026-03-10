package service

import (
	"TicketManagement/dto"
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"TicketManagement/repository"
)

type WorkflowPathService interface {
	CreateWorkflowPath(req *dto.CreateWorkflowPathRequest) error
	UpdateWorkflowPath(id int, req *dto.UpdateWorkflowPathRequest) error
	DeleteWorkflowPath(id int) error
	GetAllWorkflowPath() ([]dto.WorkflowPathResponse, error)
}

type workflowPathService struct {
	repository repository.WorkflowPathRepository
}

func NewWorkflowPathService(repository repository.WorkflowPathRepository) *workflowPathService {
	return &workflowPathService{repository: repository}
}

func (s *workflowPathService) CreateWorkflowPath(req *dto.CreateWorkflowPathRequest) error {
	wp := entity.WorkflowPath{
		WorkflowID:   req.WorkflowID,
		ParallelKey:  req.ParallelKey,
		ExeCondition: req.ExeCondition,
		ReadColumn:   req.ReadColumn,
		AssignedTo:   req.AssignedTo,
		Activity:     req.Activity,
	}

	return s.repository.CreateWorkflowPath(&wp)
}

func (s *workflowPathService) UpdateWorkflowPath(id int, req *dto.UpdateWorkflowPathRequest) error {

	wp := entity.WorkflowPath{
		ID:           id,
		ParallelKey:  req.ParallelKey,
		ExeCondition: req.ExeCondition,
		ReadColumn:   req.ReadColumn,
		AssignedTo:   req.AssignedTo,
		Activity:     req.Activity,
	}

	return s.repository.UpdateWorkflow(&wp)
}

func (s *workflowPathService) DeleteWorkflowPath(id int) error {
	return s.repository.DeleteWorkflowpath(id)
}

func (s *workflowPathService) GetAllWorkflowPath() ([]dto.WorkflowPathResponse, error) {

	list, err := s.repository.GetAllWorkflowPath()
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Data Not Found"}
	}

	var responses []dto.WorkflowPathResponse

	for _, wp := range list {

		exe := wp.ExeCondition
		read := wp.ReadColumn
		responses = append(responses, dto.WorkflowPathResponse{
			ID:           wp.ID,
			WorkflowID:   wp.WorkflowID,
			ParallelKey:  wp.ParallelKey,
			ExeCondition: exe,
			ReadColumn:   read,
			AssignedTo:   wp.AssignedTo,
			Activity:     wp.Activity,
		})
	}

	return responses, nil
}
