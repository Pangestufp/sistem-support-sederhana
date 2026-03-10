package router

import (
	"TicketManagement/config"
	"TicketManagement/handler"
	"TicketManagement/middleware"
	"TicketManagement/repository"
	"TicketManagement/service"

	"github.com/gin-gonic/gin"
)

func RoleRouter(api *gin.RouterGroup) {
	repo := repository.NewRoleRepository(config.DB)
	service := service.NewRoleService(repo, config.RedisClient)
	handler := handler.NewRoleHandler(service)

	role := api.Group("/role")
	role.Use(middleware.JWTMiddleware())

	role.POST("", middleware.RoleMiddleware(101), handler.Create)
	role.GET("", handler.GetAll)
	role.GET("/:id", handler.GetByID)
	role.PUT("/:id", middleware.RoleMiddleware(101), handler.Update)
	role.DELETE("/:id", middleware.RoleMiddleware(101), handler.Delete)
}
