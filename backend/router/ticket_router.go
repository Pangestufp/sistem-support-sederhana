package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func TicketRouter(api *gin.RouterGroup) {
	repositoryT := repository.NewTicketRepository(config.DB)
	repositoryTD := repository.NewTicketDetailRepository(config.DB)
	repositoryTA := repository.NewTicketAttachmentRepository(config.DB)
	repositoryTW := repository.NewTicketWorkflowRepository(config.DB)
	repositoryUR := repository.NewUserRoleRepository(config.DB)
	repositoryWP := repository.NewWorkflowPathRepository(config.DB)
	service := service.NewTicketService(repositoryT, repositoryTD, repositoryTA, repositoryTW, repositoryUR, repositoryWP)
	handler := handler.NewTicketHandler(service)

	tt := api.Group("/ticket")
	tt.Use(middleware.JWTMiddleware())

	tt.POST("", handler.Create)
	tt.GET("/:id", handler.GetTicket)
	tt.POST("/approve/:id", handler.ApproveAction)
	tt.POST("/reject/:id", handler.RejectAction)
	tt.POST("/return/:id", handler.ReturnAction)
	tt.POST("/review/:id", handler.ReviewAction)
	tt.GET("/getAll", handler.GetAllTickets)

	tt.GET("/details/:id", handler.GetTicketDetails)
	tt.GET("/attachments/:id", handler.GetTicketAttachments)
	tt.GET("/allJob", handler.GetUserJob)
	tt.GET("/workflow/:id", handler.GetTicketWorkflowV)

}
