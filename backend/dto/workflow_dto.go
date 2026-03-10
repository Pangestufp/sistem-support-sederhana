package dto

type CreateWorkflowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateWorkflowRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WorkflowResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
