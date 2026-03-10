package repository

import (
	"TicketManagement/entity"
	"TicketManagement/errorhandler"

	"gorm.io/gorm"
)

type WorkflowPathRepository interface {
	CreateWorkflowPath(workflowPath *entity.WorkflowPath) error
	UpdateWorkflow(workflowPath *entity.WorkflowPath) error
	DeleteWorkflowpath(id int) error
	GetAllWorkflowPath() ([]entity.WorkflowPath, error)
}

type workflowPathRepository struct {
	db *gorm.DB
}

func NewWorkflowPathRepository(db *gorm.DB) *workflowPathRepository {
	return &workflowPathRepository{
		db: db,
	}
}

func (r *workflowPathRepository) CreateWorkflowPath(workflowPath *entity.WorkflowPath) error {
	return r.db.Create(workflowPath).Error
}

func (r *workflowPathRepository) UpdateWorkflow(workflowPath *entity.WorkflowPath) error {
	result := r.db.Model(&entity.WorkflowPath{}).
		Where("id = ?", workflowPath.ID).
		Updates(map[string]interface{}{
			"workflow_id":   workflowPath.WorkflowID,
			"parallel_key":  workflowPath.ParallelKey,
			"exe_condition": workflowPath.ExeCondition,
			"read_column":   workflowPath.ReadColumn,
			"assigned_to":   workflowPath.AssignedTo,
			"activity":      workflowPath.Activity,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil
}

func (r *workflowPathRepository) DeleteWorkflowpath(id int) error {
	result := r.db.Delete(&entity.WorkflowPath{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &errorhandler.InternalServerError{Message: "No Row Effect"}
	}

	return nil

}

func (r *workflowPathRepository) GetAllWorkflowPath() ([]entity.WorkflowPath, error) {
	var workflowPath []entity.WorkflowPath

	err := r.db.Order("id ASC").Find(&workflowPath).Error
	if err != nil {
		return nil, err
	}

	return workflowPath, nil
}
