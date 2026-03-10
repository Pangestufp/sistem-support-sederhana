package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func WorkflowRouter(api *gin.RouterGroup) {
	repo := repository.NewWorkflowRepository(config.DB)
	service := service.NewWorkflowService(repo)
	handler := handler.NewWorkflowHandler(service)

	workflow := api.Group("/workflow")
	workflow.Use(middleware.JWTMiddleware())

	workflow.POST("", middleware.RoleMiddleware(101), handler.Create)
	workflow.GET("", handler.GetAll)
	workflow.GET("/:id", handler.GetByID)
	workflow.PUT("/:id", middleware.RoleMiddleware(101), handler.Update)
	workflow.DELETE("/:id", middleware.RoleMiddleware(101), handler.Delete)
}
