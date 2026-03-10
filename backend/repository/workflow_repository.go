package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"
	"errors"

	"gorm.io/gorm"
)

type WorkflowRepository interface {
	CreateWorkflow(workflow *entity.Workflow) error
	UpdateWorkflow(workflow *entity.Workflow) error
	DeleteWorkflow(id int) error
	GetWorkflowByID(id int) (*entity.Workflow, error)
	GetAllWorkflow() ([]entity.Workflow, error)
}

type workflowRepository struct {
	db *gorm.DB
}

func NewWorkflowRepository(db *gorm.DB) *workflowRepository {
	return &workflowRepository{
		db: db,
	}
}

func (r *workflowRepository) CreateWorkflow(workflow *entity.Workflow) error {
	return r.db.Create(workflow).Error
}

func (r *workflowRepository) UpdateWorkflow(workflow *entity.Workflow) error {
	result := r.db.Model(&entity.Workflow{}).
		Where("id = ?", workflow.ID).
		Updates(map[string]interface{}{
			"name":        workflow.Name,
			"description": workflow.Description,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *workflowRepository) DeleteWorkflow(id int) error {
	result := r.db.Delete(&entity.Workflow{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil

}

func (r *workflowRepository) GetWorkflowByID(id int) (*entity.Workflow, error) {
	var workflow entity.Workflow

	err := r.db.First(&workflow, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "No Row Found"}
		}
		return nil, err
	}

	return &workflow, nil
}

func (r *workflowRepository) GetAllWorkflow() ([]entity.Workflow, error) {
	var workflow []entity.Workflow

	err := r.db.Where("status = ?", 1).Order("id ASC").Find(&workflow).Error
	if err != nil {
		return nil, err
	}

	return workflow, nil
}
