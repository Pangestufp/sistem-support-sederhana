package handler

import (
	"TicketManagement/dto"
	"TicketManagement/errorhandler"
	"TicketManagement/helper"
	"TicketManagement/service"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ticketHandler struct {
	service service.TicketService
}

func NewTicketHandler(service service.TicketService) *ticketHandler {
	return &ticketHandler{
		service: service,
	}
}

func (h ticketHandler) Create(c *gin.Context) {
	var req dto.TicketRequest

	if err := c.ShouldBind(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	var filePaths []string
	uploadDir := "public/picture"
	_ = os.MkdirAll(uploadDir, 0755)

	for _, file := range req.Pictures {
		ext := filepath.Ext(file.Filename)
		newName := uuid.New().String() + ext
		dst := filepath.Join(uploadDir, newName)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			continue
		}

		filePaths = append(filePaths, dst)
	}

	id, _ := c.Get("userID")
	userID := id.(int)

	data, err := h.service.CreateTicket(req, filePaths, userID)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "success create ticket",
		Data:       data,
	}))

}

func (h *ticketHandler) ApproveAction(c *gin.Context) {

	userIDVal, ok := c.Get("userID")
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: "invalid user id",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	if err := h.service.Action(id, userID, helper.ActionApprove); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success approve document",
	}))

}

func (h *ticketHandler) RejectAction(c *gin.Context) {

	userIDVal, ok := c.Get("userID")
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: "invalid user id",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	if err := h.service.Action(id, userID, helper.ActionReject); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success reject document",
	}))

}

func (h *ticketHandler) ReviewAction(c *gin.Context) {

	var req dto.TicketReviewRequest

	if err := c.ShouldBind(&req); err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	userIDVal, ok := c.Get("userID")
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: "invalid user id",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	var filePaths []string
	uploadDir := "public/picture"
	_ = os.MkdirAll(uploadDir, 0755)

	for _, file := range req.Pictures {
		ext := filepath.Ext(file.Filename)
		newName := uuid.New().String() + ext
		dst := filepath.Join(uploadDir, newName)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			continue
		}

		filePaths = append(filePaths, dst)
	}

	if err := h.service.Review(req, filePaths, id, userID, helper.ActionReview); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success review document",
	}))

}

func (h *ticketHandler) ReturnAction(c *gin.Context) {

	userIDVal, ok := c.Get("userID")
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: "invalid user id",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	if err := h.service.Return(id, userID, helper.ActionReturn); err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success return document",
	}))

}

func (h *ticketHandler) GetUserJob(c *gin.Context) {

	userIDVal, ok := c.Get("userID")
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.UnauthorizedError{
			Message: "unauthorized",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		errorhandler.ErrorHandler(c, &errorhandler.InternalServerError{
			Message: "invalid user id",
		})
		return
	}

	data, err := h.service.GetAllJob(userID)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
	}))

}

func (h *ticketHandler) GetTicket(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	data, err := h.service.GetTicket(id)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
	}))

}

func (h *ticketHandler) GetTicketDetails(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	data, err := h.service.GetTicketDetail(id)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
	}))

}

func (h *ticketHandler) GetTicketAttachments(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	data, err := h.service.GetTicketAttachment(id)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
	}))

}

func (h *ticketHandler) GetTicketWorkflowV(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorhandler.ErrorHandler(c, &errorhandler.BadRequestError{
			Message: "invalid ticket id",
		})
		return
	}

	data, err := h.service.GetTicketWorkflowV(id)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
	}))

}

func (h *ticketHandler) GetAllTickets(c *gin.Context) {

	lastID, _ := strconv.Atoi(c.Query("last_id"))

	data, err := h.service.GetPaginateTicket(lastID)

	if err != nil {
		errorhandler.ErrorHandler(c, err)
		return
	}

	total := len(data)

	c.JSON(http.StatusOK, helper.BuildResponse(dto.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "success get data",
		Data:       data,
		Paginate: &dto.Paginate{
			LastID: data[total-1].ID,
			Total:  total,
		},
	}))

}
