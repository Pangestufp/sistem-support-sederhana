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

type roleHandler struct {
	service service.RoleService
}

func NewRoleHandler(service service.RoleService) *roleHandler {
	return &roleHandler{service: service}
}

func (h *roleHandler) Create(c *gin.Context) {
	var req dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.CreateRole(&req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "success create role",
	}))
}

func (h *roleHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid role id"})
		return
	}

	var req dto.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.UpdateRole(id, &req); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success update role",
	}))
}

func (h *roleHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid role id"})
		return
	}

	if err := h.service.DeleteRole(id); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success delete role",
	}))
}

func (h *roleHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAllRole()
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

func (h *roleHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: "invalid role id"})
		return
	}

	data, err := h.service.GetRoleByID(id)
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
