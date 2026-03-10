package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func WorkflowPathRouter(api *gin.RouterGroup) {
	repo := repository.NewWorkflowPathRepository(config.DB)
	service := service.NewWorkflowPathService(repo)
	handler := handler.NewWorkflowPathHandler(service)

	wp := api.Group("/workflow-path")
	wp.Use(middleware.JWTMiddleware())

	wp.POST("", middleware.RoleMiddleware(101), handler.Create)
	wp.GET("", handler.GetAll)
	wp.PUT("/:id", middleware.RoleMiddleware(101), handler.Update)
	wp.DELETE("/:id", middleware.RoleMiddleware(101), handler.Delete)
}
