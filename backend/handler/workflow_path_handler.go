package handler

import (
	"TicketManagement/dto"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type workflowPathHandler struct {
	service service.WorkflowPathService
}

func NewWorkflowPathHandler(service service.WorkflowPathService) *workflowPathHandler {
	return &workflowPathHandler{service: service}
}

func (h *workflowPathHandler) Create(c *gin.Context) {
	var req dto.CreateWorkflowPathRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateWorkflowPath(&req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "success create workflow path",
	}))
}

func (h *workflowPathHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateWorkflowPathRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.UpdateWorkflowPath(id, &req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success update workflow path",
	}))
}

func (h *workflowPathHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteWorkflowPath(id); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success delete workflow path",
	}))
}

func (h *workflowPathHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAllWorkflowPath()
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       data,
	}))
}
