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

type deparmentHandler struct {
	service service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *deparmentHandler {
	return &deparmentHandler{
		service: service,
	}
}

func (h *deparmentHandler) Create(c *gin.Context) {
	var req dto.CreateDepartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateDepartment(&req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "Success Create Department",
	})
	c.JSON(http.StatusCreated, res)
}

func (h *deparmentHandler) Update(c *gin.Context) {
	idParam := c.Param("id")

	departmentID, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid department id"})
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.UpdateDepartment(departmentID, &req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success update department",
	})

	c.JSON(http.StatusOK, res)
}

func (h *deparmentHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	departmentID, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid department id"})
		return
	}

	data, err := h.service.GetDepartmentByID(departmentID)
	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       data,
	})

	c.JSON(http.StatusOK, res)
}

func (h *deparmentHandler) GetAll(c *gin.Context) {

	data, err := h.service.GetallDepartment()

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(
		dto.ResponseParam{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       data,
		})

	c.JSON(http.StatusOK, res)
}

func (h *deparmentHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	departmentID, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid department id"})
		return
	}

	if err := h.service.DeleteDepartment(departmentID); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success delete department",
	})

	c.JSON(http.StatusOK, res)
}
