package handler

import (
	"TicketManagement/dto"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(
		dto.ResponseParam{
			StatusCode: http.StatusCreated,
			Message:    "Register Success",
		})

	c.JSON(http.StatusCreated, res)

}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	if err := c.ShouldBindJSON(&login); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	res := helper.BuildResponse(
		dto.ResponseParam{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       result,
		})

	c.JSON(http.StatusOK, res)
}
