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

type ticketTypeHandler struct {
	service service.TicketTypeService
}

func NewTicketTypeHandler(service service.TicketTypeService) *ticketTypeHandler {
	return &ticketTypeHandler{service: service}
}

func (h *ticketTypeHandler) Create(c *gin.Context) {
	var req dto.CreateTicketTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateTicketType(&req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "success create ticket type",
	}))
}

func (h *ticketTypeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateTicketTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.UpdateTicketType(id, &req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success update ticket type",
	}))
}

func (h *ticketTypeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.service.DeleteTicketType(id); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success delete ticket type",
	}))
}

func (h *ticketTypeHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAllTicketType()
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

func (h *ticketTypeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := h.service.GetTicketTypeByID(id)
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
