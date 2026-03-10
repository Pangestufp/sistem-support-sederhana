package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func TicketTypeRouter(api *gin.RouterGroup) {
	repo := repository.NewTicketTypeRepository(config.DB)
	service := service.NewTicketTypeService(repo, config.RedisClient)
	handler := handler.NewTicketTypeHandler(service)

	tt := api.Group("/ticket-type")
	tt.Use(middleware.JWTMiddleware())

	tt.POST("", middleware.RoleMiddleware(101), handler.Create)
	tt.GET("", handler.GetAll)
	tt.GET("/:id", handler.GetByID)
	tt.PUT("/:id", middleware.RoleMiddleware(101), handler.Update)
	tt.DELETE("/:id", middleware.RoleMiddleware(101), handler.Delete)
}
