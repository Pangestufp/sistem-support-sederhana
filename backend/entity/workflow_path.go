package entity

type WorkflowPath struct {
	ID           int
	WorkflowID   int
	ParallelKey  int
	ExeCondition string
	ReadColumn   string
	AssignedTo   string
	Activity     string
}
