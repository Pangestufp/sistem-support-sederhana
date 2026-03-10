package dto

type CreateWorkflowPathRequest struct {
	WorkflowID   int    `json:"workflow_id"`
	ParallelKey  int    `json:"parallel_key"`
	ExeCondition string `json:"exe_condition"`
	ReadColumn   string `json:"read_column"`
	AssignedTo   string `json:"assigned_to"`
	Activity     string `json:"activity"`
}

type UpdateWorkflowPathRequest struct {
	ParallelKey  int    `json:"parallel_key"`
	ExeCondition string `json:"exe_condition"`
	ReadColumn   string `json:"read_column"`
	AssignedTo   string `json:"assigned_to"`
	Activity     string `json:"activity"`
}

type WorkflowPathResponse struct {
	ID           int    `json:"id"`
	WorkflowID   int    `json:"workflow_id"`
	ParallelKey  int    `json:"parallel_key"`
	ExeCondition string `json:"exe_condition"`
	ReadColumn   string `json:"read_column"`
	AssignedTo   string `json:"assigned_to"`
	Activity     string `json:"activity"`
}
